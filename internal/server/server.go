package server

import (
	"net/http"

	"github.com/dannyoka/go-server/internal/controllers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitRouter(client *mongo.Client) {
	router := mux.NewRouter().StrictSlash(true)
	router = controllers.InitApiController(router, client)
	log.SetFormatter(&log.JSONFormatter{})
	log.Info("hello there!")
	log.Info("hi again!")
	router.PathPrefix("/").Subrouter()
	http.ListenAndServe(":3000", router)
}
