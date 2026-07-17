package models

import "net/http"

type ResponseRecorder struct {
	Writer http.ResponseWriter
	Status int
}

func (r *ResponseRecorder) Header() http.Header {
	return r.Writer.Header()
}

func (r *ResponseRecorder) Write(bytes []byte) (int, error) {
	return r.Writer.Write(bytes)
}

func (r *ResponseRecorder) WriteHeader(status int) {
	r.Status = status
	r.Writer.WriteHeader(status)
}
