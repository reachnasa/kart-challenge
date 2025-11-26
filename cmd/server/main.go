package main

import (
	"food-app/internal/repository"
	"food-app/internal/routes"
	"log"
	"net/http"
)

func main() {

	repository.InitSampleData()

	router := routes.SetupRouter()

	log.Println("Server Started on")
	log.Println("\n Food App endpoints:")
	log.Println("  GET  /api/product           - List all products")
	log.Println("  GET  /api/product/{id}      - Get product by ID")
	log.Println("  POST /api/order             - Place order (requires api_key header)")
	log.Println("  GET  /health                - Health check")

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}

}
