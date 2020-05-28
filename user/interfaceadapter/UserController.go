package interfaceadapter

import (
	"example/user/usecase"
)

// UserController is a controller that controls users
type UserController struct {
}

// CreateUser is responsible for creating users
func (user UserController) CreateUser(username string) UserViewModel {
	interactor := usecase.CreateUserInteractor{}
	presenter := UserPresenter{}
	input := usecase.CreateUserInputData{Username: username}
	interactor.Handle(input, presenter)
	return presenter.ViewModel
}
