package handler

import (
	_ "github.com/k0kubun/pp"
	"github.com/zenazn/goji/web"
	"net/http"
)

type SimpleHandler struct{}

func (s *SimpleHandler) Top(c web.C, w http.ResponseWriter, r *http.Request) {
	//TODO
}

func (s *SimpleHandler) Entry(c web.C, w http.ResponseWriter, r *http.Request) {
	//TODO
}

func (s *SimpleHandler) New(c web.C, w http.ResponseWriter, r *http.Request) {
	//TODO
}

func (s *SimpleHandler) Create(c web.C, w http.ResponseWriter, r *http.Request) {
	//TODO
}
