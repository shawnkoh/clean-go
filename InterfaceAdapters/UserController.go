package interfaceadapters

import "../usecases"

// UserController is a controller that controls users
type UserController struct {
}

// CreateUser is responsible for creating users
func (user UserController) CreateUser(username string) {
	interactor := usecases.CreateUserInteractor{}
	presenter := UserPresenter{}
	input := usecases.CreateUserInputData{Username: "abcdefg"}
	interactor.Handle(input, presenter)
}
