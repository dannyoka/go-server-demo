package controllers

import (
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitApiController(router *mux.Router, client *mongo.Client) *mux.Router {
	todosController := NewTodosController(client)
	peopleController := NewPeopleController(client)
	itemsController := NewItemsController(client)
	subrouter := router.PathPrefix("/api").Subrouter()
	subrouter = peopleController.Init(subrouter)
	subrouter = todosController.Init(subrouter)
	subrouter = itemsController.Init(subrouter)
	return router
}