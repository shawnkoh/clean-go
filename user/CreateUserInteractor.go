package user

// CreateUserInputBoundary represents an interactor that creates a User
type CreateUserInputBoundary interface {
	Handle(input CreateUserInputData, presenter CreateUserOutputBoundary)
}

// CreateUserInputData represents the data required for creating a User
type CreateUserInputData struct {
	Username string
}

// CreateUserOutputData represents the data returned from creating a User
type CreateUserOutputData struct {
	ID       string
	Username string
}

// CreateUserOutputBoundary represents a presenter who consumes an output
type CreateUserOutputBoundary interface {
	Present(output CreateUserOutputData)
}

type createUserInteractor struct {
	repository Repository
}

// Handle the CreateUserInputData by updating the entities
func (interactor createUserInteractor) Handle(input CreateUserInputData, presenter CreateUserOutputBoundary) {
	// TODO: Make repository throwable
	user := interactor.repository.CreateUser(input.Username)
	output := CreateUserOutputData{ID: user.ID, Username: user.Username}
	presenter.Present(output)
}

// NewCreateUserInteractor creates a new CreateUserInteractor
func NewCreateUserInteractor(repository Repository) CreateUserInputBoundary {
	return &createUserInteractor{repository: repository}
}
