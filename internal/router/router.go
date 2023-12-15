package router

import (
	"github.com/gorilla/mux"
	"golang/internal/handlers"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/resource", handlers.GetResource).Methods("GET")
	r.HandleFunc("/api/resource", handlers.CreateResource).Methods("POST")
	// Adicione mais rotas conforme necessário

	return r
}
