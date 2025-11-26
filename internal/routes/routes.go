package routes

import (
	"food-app/internal/filter"
	"food-app/internal/handler"
	"food-app/internal/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	r.Use(filter.Logging)

	// Initialize controllers
	productController := handler.NewProductController()
	orderController := handler.NewOrderController()

	// API routes under /api prefix
	api := r.PathPrefix("/api").Subrouter()

	// Product endpoints (no auth required)
	api.HandleFunc("/product", productController.ListProducts).Methods("GET")
	api.HandleFunc("/product/{productId}", productController.GetProduct).Methods("GET")

	// Order endpoints (require API key)
	orderRouter := api.PathPrefix("/order").Subrouter()
	orderRouter.Use(filter.ApiKey)
	orderRouter.HandleFunc("", orderController.PlaceOrder).Methods("POST")

	// Health check endpoint
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		utils.RespondJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	}).Methods("GET")

	return r
}
