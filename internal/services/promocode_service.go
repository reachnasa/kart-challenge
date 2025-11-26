package services

import (
	"bufio"
	"context"
	"log"
	"os"
	"strings"
	"sync"
)

type PromoCodeService struct {
	promoFiles []string
}

func NewPromoCodeService() *PromoCodeService {
	return &PromoCodeService{
		promoFiles: []string{
			"../../data/couponbase1",
			"../../data/couponbase2",
			"../../data/couponbase3",
		},
	}
}

// ValidatePromoCode searches for coupon code
func (s *PromoCodeService) ValidatePromoCode(code string) (bool, error) {
	if code == "" {
		return false, nil
	}

	// Rule 1: Validate length (8-10 characters)
	codeLen := len(code)
	if codeLen < 8 || codeLen > 10 {
		log.Printf("Promo code '%s' invalid: length is %d, must be 8-10 characters", code, codeLen)
		return false, nil
	}

	log.Printf("Promo code '%s' length validation passed (%d characters)", code, codeLen)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Channel to receive promo code results from goroutines
	resultChan := make(chan string, len(s.promoFiles))
	var wg sync.WaitGroup

	// Launch a goroutine for each coupon file
	for _, filePath := range s.promoFiles {
		wg.Add(1)
		go s.searchPromoCodeInFile(ctx, filePath, code, resultChan, &wg)
	}

	// Goroutine to close result channel when all searches complete
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Collect results from channel
	var foundFiles []string
	foundCount := 0

	for result := range resultChan {
		foundCount++

		log.Printf("Promo code '%s' found in %s [%d/2]",
			code, result, foundCount)
		foundFiles = append(foundFiles, result)

		// Rule 2: Cancel all goroutines when found in 2 files
		if foundCount >= 2 {
			log.Printf("Promo code '%s' VALIDATED! Found in %d files: %v",
				code, foundCount, foundFiles)
			log.Printf("COUPON APPLIED: %s", code)
			cancel()
			break
		}
	}

	// Wait for all goroutines to finish
	wg.Wait()

	if foundCount >= 2 {
		return true, nil
	}

	if foundCount == 1 {
		log.Printf("Promo code '%s' INVALID: found in only 1 file (%s). Must be in at least 2 files.",
			code, foundFiles[0])
	} else {
		log.Printf("Promo code '%s' INVALID: not found in any files.", code)
	}

	return false, nil
}

// searchPromoCodeInFile searches for promo code in a single file using streaming
func (s *PromoCodeService) searchPromoCodeInFile(
	ctx context.Context,
	filePath string,
	searchCode string,
	resultChan chan<- string,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	file, err := os.Open(filePath)
	if err != nil {
		log.Printf("Warning: Could not open %s: %v", filePath, err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 0

	// Stream file line by line
	for scanner.Scan() {
		// Check if context is cancelled
		select {
		case <-ctx.Done():
			log.Printf("Search cancelled in %s", filePath)
			return
		default:
		}

		lineNum++
		line := strings.TrimSpace(scanner.Text())

		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		code := strings.TrimSpace(line)

		// Check if this is the code is valid
		if strings.EqualFold(code, searchCode) {
			// Send result to channel

			select {
			case resultChan <- filePath:
				return // Code Found, stop searching this file
			case <-ctx.Done():
				return // Context cancelled
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error reading %s: %v", filePath, err)
	}
}
