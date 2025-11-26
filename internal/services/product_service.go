package services

import (
	"food-app/internal/models"
	"food-app/internal/repository"
	"sort"
	"strconv"
)

type ProductService struct{}

func NewProductService() *ProductService {
	return &ProductService{}
}

func (s *ProductService) GetAllProducts() []models.Product {

	productList := repository.GetAllProducts()
	/* sort.Slice(productList, func(i, j int) bool {
		return productList[i].ID < productList[j].ID
	}) */

	sort.Slice(productList, func(i, j int) bool {
		id1, _ := strconv.Atoi(productList[i].ID)
		id2, _ := strconv.Atoi(productList[j].ID)
		return id1 < id2
	})

	return productList
}

func (s *ProductService) GetProductByID(id string) (*models.Product, error) {
	return repository.GetProductByID(id)
}
