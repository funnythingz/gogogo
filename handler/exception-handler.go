package handler

import (
	"net/http"
)

type ExceptionHandler struct{}

func (_ *ExceptionHandler) NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}
