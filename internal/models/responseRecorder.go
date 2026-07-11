package models

import "net/http"

type ResponseRecorder struct {
	Writer http.ResponseWriter
	Status int
}

func (r *ResponseRecorder) WriteHeader(status int) {
	r.Status = status
	r.Writer.WriteHeader(status)
}
