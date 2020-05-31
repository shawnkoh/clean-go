package user

// Controller is a controller that controls users
type Controller interface {
	CreateUser(username string)
}

type controller struct {
	interactor CreateUserInputBoundary
	presenter  CreateUserOutputBoundary
}

// CreateUser is responsible for creating users
func (controller controller) CreateUser(username string) {
	input := CreateUserInputData{Username: username}
	controller.interactor.Handle(input, controller.presenter)
}

// NewController creates a UserController
func NewController(interactor CreateUserInputBoundary, presenter CreateUserOutputBoundary) Controller {
	return &controller{interactor: interactor, presenter: presenter}
}
