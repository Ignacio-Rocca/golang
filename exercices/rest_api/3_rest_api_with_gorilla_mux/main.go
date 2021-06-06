package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "get called"}`))
}

func post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "post called"}`))
}

func put(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(`{"message": "put called"}`))
}

func delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "delete called"}`))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "not found"}`))
}

func main() {
	r := mux.NewRouter()

	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("", get).Methods(http.MethodGet)
	api.HandleFunc("", post).Methods(http.MethodPost)
	api.HandleFunc("", put).Methods(http.MethodPut)
	api.HandleFunc("", delete).Methods(http.MethodDelete)
	api.HandleFunc("", notFound)

	usersRouter := api.PathPrefix("/users").Subrouter()
	usersRouter.HandleFunc("", getUsers).Methods("GET")
	usersRouter.HandleFunc("/{user_id}", getUserByID).Methods("GET")

	log.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

type User struct {
	ID   string
	Name string
}

var UsersRepository = []User{
	{
		ID:   "1",
		Name: "Ignacio Rocca",
	},
	{
		ID:   "2",
		Name: "Diego Maradona",
	},
	{
		ID:   "3",
		Name: "Lionel Messi",
	},
}

// example handling query params
func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	queryParams := r.URL.Query()
	limit := queryParams.Get("limit")
	limitInt := 0
	if(strings.Trim(limit, " ") != "") {
		limit, err := strconv.Atoi(limit)
		if err != nil || limit <= 0 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"message":"invalid limit query param"}`))
		}
		limitInt = limit
	}

	var usersResponse []User
	if limitInt != 0 && limitInt < len(UsersRepository){
		for i := 0; i < limitInt; i++ {
			usersResponse = append(usersResponse, UsersRepository[i])
		}
	} else {
		usersResponse = UsersRepository
	}

	w.WriteHeader(http.StatusOK)
	response, err := json.Marshal(usersResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(response)
	return
}

// example handling path params
func getUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	pathParams := mux.Vars(r)
	userID, ok := pathParams["user_id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"invalid user_id"}`))
		return
	}

	for _, u := range UsersRepository {
		if u.ID == userID {
			w.WriteHeader(http.StatusOK)
			userJson, err := json.Marshal(u)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write(userJson)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"user not found"}`))
	return
}

