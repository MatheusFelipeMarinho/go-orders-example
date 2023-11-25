package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github/braip.com/internal/application/service"
	"github/braip.com/internal/domain/entity"
	"net/http"
	"strconv"
)

type OrderController struct {
	orderService *service.OrderService
}

func NewOrderController(orderService *service.OrderService) *OrderController {
	return &OrderController{
		orderService: orderService,
	}
}

func (oc *OrderController) HandleListOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := oc.orderService.ListOrder()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(orders)
}

func (oc *OrderController) HandleGetOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}
	order, err := oc.orderService.GetOrder(uint(orderID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(order)
}

func (oc *OrderController) HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
	var order entity.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := oc.orderService.CreateOrder(&order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(order)
}
