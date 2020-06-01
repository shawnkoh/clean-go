package main

import (
	"encoding/json"
	"example/database"
	"example/user"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/users", createUser)
	http.ListenAndServe(":3000", r)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	// get input from request
	repository := database.UserRepository{}
	controller := user.NewController(repository)
	viewModel := controller.CreateUser("tommy")

	// reply to client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	bytes, _ := json.Marshal(&viewModel)
	w.Write(bytes)
}
