package handlers

import (
	"net/http"
)

type Resource struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	// Adicione mais campos conforme necessário
}

func GetResource(w http.ResponseWriter, r *http.Request) {
	// Lógica para obter recursos e responder com JSON
}

func CreateResource(w http.ResponseWriter, r *http.Request) {
	// Lógica para criar um novo recurso a partir do corpo JSON da solicitação
}
