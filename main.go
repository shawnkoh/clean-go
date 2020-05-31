package main

import (
	"errors"
	"example/database"
	"example/user"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

func main() {
	r := chi.NewRouter()

	// Middlewares
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("welcome"))
	})

	r.Route("/users", func(r chi.Router) {
		r.Post("/", createUser)
	})

	http.ListenAndServe(":3000", r)
}

// CreateUserRequest represents a POST request to create a User
type CreateUserRequest struct {
	username string
}

// Bind conforms CreateUserRequest to render.Binder interface
func (input CreateUserRequest) Bind(r *http.Request) error {
	if input.username == "" {
		return errors.New("missing field")
	}

	return nil
}

func createUser(w http.ResponseWriter, r *http.Request) {
	data := &CreateUserRequest{}
	render.Bind(r, data)

	repository := database.UserRepository{}
	interactor := user.NewCreateUserInteractor(repository)
	presenter := user.Presenter{}
	controller := user.NewController(interactor, presenter)

	controller.CreateUser(data.username)
	viewModel := presenter.ViewModel

	render.Status(r, http.StatusCreated)
	render.Render(w, r, CreateUserResponse{ID: viewModel.ID, Username: viewModel.Username})
}

// CreateUserResponse is the response given to the caller
type CreateUserResponse struct {
	ID       string
	Username string
}

// Render conforms CreateUserResponse to render.Render interface
func (response CreateUserResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
