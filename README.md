# Food API Server

A Go REST API to manage products and place orders with promo code validation.

---

## Features

- Health check: curl http://localhost:8080/health
- List all products: `/api/product`  
- Get product by ID: `/api/product/{id}`  
- Place an order: `/api/order` with promo codes  
- Promo code rules:
  - 8-10 characters long  
  - Must appear in at least 2 of 3 promo files (`data/couponbase1`, `couponbase2`, `couponbase3`)  

---

## Run

```bash
git clone https://github.com/reachnasa/kart-challenge.git
cd cmd/server
go mod tidy
go run main.go

Server runs at: http://localhost:8080
```
---

## Notes
- Uses goroutines & channels for concurrent promo code validation
- Stops searching once a promo code is found in 2 files
- Designed to handle large promo files efficiently
- Large coupon data files (`couponbase1`, `couponbase2`, `couponbase3`) are stored using **Git LFS** due to their size