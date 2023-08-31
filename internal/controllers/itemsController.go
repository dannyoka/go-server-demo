package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dannyoka/go-server/internal/types"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)


var items = []types.Todo{{Name: "Go to the store"}, {Name: "Pick up the kids"}}

type IItemsController interface {
	getItemsHandler(w http.ResponseWriter, r *http.Request)
	Init(router *mux.Router )  *mux.Router
}

type ItemsController struct {
	Client *mongo.Client
}

func NewItemsController(client *mongo.Client) IItemsController {
	return &ItemsController{
		Client: client,
	}
}

func (c *ItemsController) getItemsHandler(w http.ResponseWriter, r *http.Request){
	response, err := json.Marshal(items)
	if err != nil {
		fmt.Println("error marshaling items")
		w.WriteHeader(500)
		w.Write([]byte("Internal server error"))
		return
	}
	w.Write(response)
}

func (c *ItemsController)Init (router *mux.Router) *mux.Router {
	subrouter := router.PathPrefix("/items").Subrouter()
	subrouter.HandleFunc("/", c.getItemsHandler).Methods(http.MethodGet)
	return router
}