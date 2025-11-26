# Food API Server

A Go REST API to manage products and place orders with promo code validation.

---

## Features

- List all products: `/api/product`  
- Get product by ID: `/api/product/{id}`  
- Place an order: `/api/order` with promo codes  
- Promo code rules:
  - 8-10 characters long  
  - Must appear in at least 2 of 3 promo files (`data/couponbase1`, `couponbase2`, `couponbase3`)  

---

## Run

```bash
git clone https://github.com/your-username/food-api.git
cd food-api
go mod tidy
go run main.go

Server runs at: http://localhost:8080

API Examples

List products
curl http://localhost:8080/api/product

Get product by ID
curl http://localhost:8080/api/product/1

Place an order with promo codes
curl -X POST http://localhost:8080/api/order \
-H "Content-Type: application/json" \
-H "api_key: apitest" \
-d '{
  "productId": "1",
  "quantity": 2,
  "promoCode": "HAPPYHRS"
}'

Health check
curl http://localhost:8080/health


Notes

Uses goroutines & channels for concurrent promo code validation

Stops searching once a promo code is found in 2 files

Designed to handle large promo files efficiently