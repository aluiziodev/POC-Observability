package handler

import (
	"encoding/json"
	"net/http"
	"order-service/internal/models"
	"order-service/internal/response"
	"order-service/internal/service"

	"github.com/gorilla/mux"
)

func Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req models.CreateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, "Erro no corpo da requisiçao")
		return
	}

	order, err := service.CreateOrder(ctx, models.CreateOrderInput{
		Item:     req.Item,
		Quantity: req.Quantity,
	})
	if err != nil {
		response.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.WriteJSON(w, http.StatusCreated, order)

}

func List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	orders, err := service.ListOrders(ctx)
	if err != nil {
		response.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.WriteJSON(w, http.StatusFound, orders)
}

func Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	parameters := mux.Vars(r)

	ID := parameters["id"]

	order, err := service.GetOrder(ctx, ID)
	if err != nil {
		response.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.WriteJSON(w, http.StatusFound, order)

}

/*func Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	parameters := mux.Vars(r)

	ID := parameters["id"]


}
*/
