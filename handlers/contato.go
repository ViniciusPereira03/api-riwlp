package handlers

import (
	"api/models"
	"api/services"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

// HandlerContext contém o cliente MongoDB
type HandlerContext struct {
	Client *mongo.Client
}

// NewHandlerContext cria um novo HandlerContext
func NewHandlerContext(client *mongo.Client) *HandlerContext {
	return &HandlerContext{Client: client}
}

// AddContato é o manipulador para adicionar contatos
func (ctx *HandlerContext) AddContato(w http.ResponseWriter, r *http.Request) {
	// Limitar o tamanho do formulário analisado (em bytes)
	// err := r.ParseMultipartForm(10 << 20) // 10 MB
	// if err != nil {
	// 	http.Error(w, "Erro ao analisar o formulário", http.StatusBadRequest)
	// 	return
	// }

	var data models.Contato
	data.Name = r.FormValue("name")
	data.Email = r.FormValue("email")
	data.Message = r.FormValue("message")

	contato := services.AddContato(ctx.Client, data)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contato)
}
