package handler

import (
	"encoding/json"
	"food-app/internal/models"
	"food-app/internal/services"
	"food-app/internal/utils"
	"net/http"
)

type OrderController struct {
	service *services.OrderService
}

func NewOrderController() *OrderController {
	return &OrderController{
		service: services.NewOrderService(),
	}
}

func (c *OrderController) PlaceOrder(w http.ResponseWriter, r *http.Request) {
	var orderReq models.OrderReq
	if err := json.NewDecoder(r.Body).Decode(&orderReq); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	order, err := c.service.PlaceOrder(orderReq)
	if err != nil {
		statusCode := http.StatusUnprocessableEntity
		if err.Error() == "order must contain at least one item" {
			statusCode = http.StatusBadRequest
		}
		utils.RespondError(w, statusCode, err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusOK, order)
}
