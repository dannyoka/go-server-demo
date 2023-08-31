package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dannyoka/go-server/internal/types"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)


var todos = []types.Todo{{Name: "Go to the store"}, {Name: "Pick up the kids"}}

type ITodosController interface {
	getTodosHandler(w http.ResponseWriter, r *http.Request)
	Init(router *mux.Router) *mux.Router
}

type TodosController struct {
	Client *mongo.Client
}

func NewTodosController(client *mongo.Client) ITodosController{
	return &TodosController{Client: client}
}

func (c *TodosController)getTodosHandler(w http.ResponseWriter, r *http.Request){
	response, err := json.Marshal(todos)
	if err != nil {
		fmt.Println("error marshaling todos")
		w.WriteHeader(500)
		w.Write([]byte("Internal server error"))
		return
	}
	w.Write(response)
}

func (c *TodosController)Init (router *mux.Router) *mux.Router {
	subrouter := router.PathPrefix("/todos").Subrouter()
	subrouter.HandleFunc("/", c.getTodosHandler).Methods(http.MethodGet)
	return router
}