package handler

import (
	"encoding/json"
	"net/http"
	"order-service/internal/domain"
	"order-service/internal/models"
	"order-service/internal/observability"
	"order-service/internal/response"
	"order-service/internal/service"
)

// POST - /orders

func Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req models.CreateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		observability.WarnContext(ctx, "Corpo da requisiçao invalido!", err)
		response.ErrorResponse(w, http.StatusBadRequest, "Erro no corpo da requisiçao")
		return
	}

	order, err := service.CreateOrder(ctx, models.CreateOrderInput{
		Item:     req.Item,
		Quantity: req.Quantity,
	})
	if err != nil {
		observability.ErrorContext(ctx, "Erro na criaçao do pedido!", err)
		response.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.WriteJSON(w, http.StatusCreated, order)

}

// GET - /orders

func List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	orders, err := service.ListOrders(ctx)
	if err != nil {
		observability.ErrorContext(ctx, "Erro ao listar os pedidos!", err)
		response.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.WriteJSON(w, http.StatusOK, orders)
}

// GET - /orders/{id}

func Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ID := r.PathValue("id")

	order, err := service.GetOrder(ctx, ID)
	if err != nil {
		observability.ErrorContext(ctx, "Erro ao listar pedido!", err)
		response.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.WriteJSON(w, http.StatusOK, order)

}

// PATCH - /orders/{id}/status

func Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ID := r.PathValue("id")

	var req models.StatusRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		observability.WarnContext(ctx, "Corpo da requisiçao invalido!", err)
		response.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	order, err := service.UpdateOrderStatus(ctx, ID, domain.OrderStatus(req.Status))
	if err != nil {
		observability.ErrorContext(ctx, "Erro na atualizaçao do pedido!", err)
		response.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.WriteJSON(w, http.StatusOK, order)

}
