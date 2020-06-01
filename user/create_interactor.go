package user

// CreateUserInputBoundary represents an interactor that creates a User
type CreateUserInputBoundary interface {
	Handle(input createUserInputData, presenter createUserOutputBoundary)
}

// CreateUserInputData represents the data required for creating a User
type createUserInputData struct {
	username string
}

// CreateUserOutputData represents the data returned from creating a User
type createUserOutputData struct {
	id       string
	username string
}

// createUserOutputBoundary represents a presenter who consumes an output
type createUserOutputBoundary interface {
	Present(output createUserOutputData)
}

type createUserInteractor struct {
	repository Repository
}

// Handle the CreateUserInputData by updating the entities
func (interactor createUserInteractor) Handle(input createUserInputData, presenter createUserOutputBoundary) {
	// TODO: Make repository throwable
	user := interactor.repository.CreateUser(input.username)
	output := createUserOutputData{id: user.ID, username: user.Username}
	presenter.Present(output)
}

// NewCreateUserInteractor creates a new CreateUserInteractor
func NewCreateUserInteractor(repository Repository) CreateUserInputBoundary {
	return &createUserInteractor{repository: repository}
}
