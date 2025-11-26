package handler

import (
	"food-app/internal/services"
	"food-app/internal/utils"
	"net/http"

	"github.com/gorilla/mux"
)

type ProductController struct {
	service *services.ProductService
}

func NewProductController() *ProductController {
	return &ProductController{
		service: services.NewProductService(),
	}
}

func (c *ProductController) ListProducts(w http.ResponseWriter, r *http.Request) {
	products := c.service.GetAllProducts()
	utils.RespondJSON(w, http.StatusOK, products)
}

func (c *ProductController) GetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID := vars["productId"]

	product, err := c.service.GetProductByID(productID)
	if err != nil {
		utils.RespondError(w, http.StatusNotFound, "Product not found")
		return
	}

	utils.RespondJSON(w, http.StatusOK, product)
}
