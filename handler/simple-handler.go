package handler

import (
	_ "github.com/k0kubun/pp"
	"github.com/zenazn/goji/web"
	"net/http"
)

type SimpleHandler struct{}

// GET
func (h *SimpleHandler) Top(c web.C, w http.ResponseWriter, r *http.Request) {
	//TODO
}

// GET
func (h *SimpleHandler) Entry(c web.C, w http.ResponseWriter, r *http.Request) {
	//TODO
}

// POST
func (h *SimpleHandler) Create(c web.C, w http.ResponseWriter, r *http.Request) {
	//TODO
}
