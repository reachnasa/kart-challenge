package services

import (
	"errors"
	"food-app/internal/models"
	"food-app/internal/repository"
	"log"

	"github.com/google/uuid"
)

type OrderService struct {
	promoCodeService *PromoCodeService
}

func NewOrderService() *OrderService {
	return &OrderService{
		promoCodeService: NewPromoCodeService(),
	}
}

func (s *OrderService) PlaceOrder(orderReq models.OrderReq) (*models.Order, error) {
	// Validate order has items
	if len(orderReq.Items) == 0 {
		return nil, errors.New("order must contain at least one item")
	}

	// Validate all items and collect products
	var products []models.Product
	for _, item := range orderReq.Items {
		if item.Quantity <= 0 {
			return nil, errors.New("item quantity must be greater than 0")
		}

		product, err := repository.GetProductByID(item.ProductID)
		if err != nil {
			return nil, errors.New("product not found: " + item.ProductID)
		}
		products = append(products, *product)
	}

	// Create order
	order := &models.Order{
		ID:       uuid.New().String(),
		Items:    orderReq.Items,
		Products: products,
	}

	// Validate promo code concurrently if provided
	if orderReq.CouponCode != "" {
		log.Printf("Validating promo code '%s' across multiple files...", orderReq.CouponCode)

		isValid, err := s.promoCodeService.ValidatePromoCode(orderReq.CouponCode)
		if err != nil {
			return nil, errors.New("error validating promo code: " + err.Error())
		}

		if !isValid {
			log.Printf("Order placed without discount - invalid promo code")
		}
	}

	// Save order
	if err := repository.CreateOrder(order); err != nil {
		return nil, err
	}

	return order, nil
}
