package repository

import (
	"errors"
	"food-app/internal/models"
	"sync"
)

type Storage struct {
	mu       sync.RWMutex
	products map[string]*models.Product
	orders   map[string]*models.Order
}

var store = &Storage{
	products: make(map[string]*models.Product),
	orders:   make(map[string]*models.Order),
}

func GetAllProducts() []models.Product {
	store.mu.RLock()
	defer store.mu.RUnlock()

	products := make([]models.Product, 0, len(store.products))
	for _, p := range store.products {
		products = append(products, *p)
	}
	return products
}

func GetProductByID(id string) (*models.Product, error) {
	store.mu.RLock()
	defer store.mu.RUnlock()

	product, exists := store.products[id]
	if !exists {
		return nil, errors.New("product not found")
	}
	return product, nil
}

func CreateOrder(order *models.Order) error {
	store.mu.Lock()
	defer store.mu.Unlock()

	store.orders[order.ID] = order
	return nil
}

func GetOrderByID(id string) (*models.Order, error) {
	store.mu.RLock()
	defer store.mu.RUnlock()

	order, exists := store.orders[id]
	if !exists {
		return nil, errors.New("order not found")
	}
	return order, nil
}

// Initialize sample data
func InitSampleData() {
	sampleProducts := []*models.Product{
		{ID: "1", Name: "Chicken Waffle", Price: 12.5, Category: "Waffle"},
		{ID: "2", Name: "Belgian Waffle", Price: 10.5, Category: "Waffle"},
		{ID: "3", Name: "Pancake Stack", Price: 8, Category: "Pancake"},
		{ID: "4", Name: "Blueberry Pancakes", Price: 9, Category: "Pancake"},
		{ID: "5", Name: "French Toast", Price: 11.5, Category: "Toast"},
		{ID: "6", Name: "Breakfast Burrito", Price: 13, Category: "Burrito"},
		{ID: "7", Name: "Eggs Benedict", Price: 14, Category: "Eggs"},
		{ID: "8", Name: "Avocado Toast", Price: 10.5, Category: "Toast"},
		{ID: "9", Name: "Caesar Salad", Price: 9, Category: "Salad"},
		{ID: "10", Name: "Chicken Salad", Price: 9, Category: "Salad"},
	}

	store.mu.Lock()
	defer store.mu.Unlock()
	for _, p := range sampleProducts {
		store.products[p.ID] = p
	}
}
