package handler

import (
	"net/http"
	"order-service/internal/response"
)

func Health(w http.ResponseWriter, r *http.Request) {
	response.WriteJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}
