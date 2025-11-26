package services

import (
	"os"
	"testing"
)

// create a temporary file with given lines
func createTempFile(t *testing.T, dir string, lines []string) string {
	file, err := os.CreateTemp(dir, "promo*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}

	for _, line := range lines {
		file.WriteString(line + "\n")
	}
	file.Close()

	return file.Name()
}

// --- TESTS ---

func TestPromoCodeService_ValidatePromoCode(t *testing.T) {

	tempDir := t.TempDir()

	// TEST DATA
	// Two files contain the code -> should return VALID
	file1 := createTempFile(t, tempDir, []string{"ABC12345"})
	file2 := createTempFile(t, tempDir, []string{"ABC12345", "XYZ99999"})
	file3 := createTempFile(t, tempDir, []string{"NOTHING"})

	// Create service with our temp files
	service := &PromoCodeService{
		promoFiles: []string{file1, file2, file3},
	}

	valid, err := service.ValidatePromoCode("ABC12345")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if !valid {
		t.Errorf("Expected promo code to be valid (found in 2 files), got invalid")
	}
}

func TestPromoCodeService_Invalid_NotEnoughMatches(t *testing.T) {
	tempDir := t.TempDir()

	// Only 1 file contains the code
	file1 := createTempFile(t, tempDir, []string{"ONLYONEE"})
	file2 := createTempFile(t, tempDir, []string{"NOTHING"})
	file3 := createTempFile(t, tempDir, []string{"NOTHING"})

	service := &PromoCodeService{
		promoFiles: []string{file1, file2, file3},
	}

	valid, _ := service.ValidatePromoCode("ONLYONEE")
	if valid {
		t.Errorf("Expected INVALID (only 1 file matches), got VALID")
	}
}

func TestPromoCodeService_Invalid_NotFound(t *testing.T) {
	tempDir := t.TempDir()

	file1 := createTempFile(t, tempDir, []string{"AAAAAAA"})
	file2 := createTempFile(t, tempDir, []string{"BBBBBBBB"})
	file3 := createTempFile(t, tempDir, []string{"CCCCCCC"})

	service := &PromoCodeService{
		promoFiles: []string{file1, file2, file3},
	}

	valid, _ := service.ValidatePromoCode("NOPE1234")
	if valid {
		t.Errorf("Expected promo code to be invalid (not found), got VALID")
	}
}

func TestPromoCodeService_Invalid_Length(t *testing.T) {
	service := &PromoCodeService{}

	valid, _ := service.ValidatePromoCode("123") // too short
	if valid {
		t.Errorf("Expected INVALID for short code, got VALID")
	}
}
