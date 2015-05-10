package handler

import (
	_ "github.com/k0kubun/pp"
	"github.com/zenazn/goji/web"
	"net/http"
)

type SimpleHandler struct{}

func (h *SimpleHandler) Top(c web.C, w http.ResponseWriter, r *http.Request) {
	//TODO
}

func (h *SimpleHandler) Entry(c web.C, w http.ResponseWriter, r *http.Request) {
	//TODO
}

func (h *SimpleHandler) New(c web.C, w http.ResponseWriter, r *http.Request) {
	//TODO
}

func (h *SimpleHandler) Create(c web.C, w http.ResponseWriter, r *http.Request) {
	//TODO
}
