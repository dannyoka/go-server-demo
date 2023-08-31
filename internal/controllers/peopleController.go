package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dannyoka/go-server/internal/types"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

var people = []types.Person{{Name: "Danny", Age: "33"}}

type IPeopleController interface {
	GetPeopleHandler(w http.ResponseWriter, r *http.Request)
	GetPersonHandler(w http.ResponseWriter, r *http.Request)
	PostPeopleHandler(w http.ResponseWriter, r *http.Request)
	DeletePeopleHandler(w http.ResponseWriter, r *http.Request)
	Init(router *mux.Router) *mux.Router
}

type PeopleController struct {
	Client *mongo.Client
}

func NewPeopleController(client *mongo.Client) IPeopleController {
	return &PeopleController{
		Client: client,
	}
}

func (c *PeopleController)GetPeopleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get people handler called")
	response, err := json.Marshal(people)
	if err != nil {
		w.Write([]byte("People not found"))
	}
	w.Write(response)
}

func (c *PeopleController)GetPersonHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	intId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("not a valid id")
		w.Write([]byte("not a valid id"))
		return
	}
	fmt.Println("Get person handler called")
	response, err := json.Marshal(people[intId])
	if err != nil {
		w.Write([]byte("People not found"))
		return
	}
	w.Write(response)
}

func (c *PeopleController)PostPeopleHandler(w http.ResponseWriter, r *http.Request) {
	var person types.Person
	json.NewDecoder(r.Body).Decode(&person)
	fmt.Println("Post people handler called")
	people = append(people, person)
	response, _ := json.Marshal(people)
	w.Write(response)
}

func (c *PeopleController)DeletePeopleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete people called")
	w.Write([]byte("hello there!"))
}

func (c *PeopleController)Init(router *mux.Router) *mux.Router{
	subrouter := router.PathPrefix("/people").Subrouter()
	subrouter.HandleFunc("/", c.GetPeopleHandler).Methods("GET")
	subrouter.HandleFunc("/{id}", c.GetPersonHandler).Methods("GET")
	subrouter.HandleFunc("/", c.PostPeopleHandler).Methods("POST")
	subrouter.HandleFunc("/", c.DeletePeopleHandler).Methods("DELETE")

	return router
}
