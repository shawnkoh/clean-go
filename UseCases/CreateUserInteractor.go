package usecases

// CreateUserInputBoundary represents an interactor that creates a User
type CreateUserInputBoundary interface {
	handle(input CreateUserInputData, presenter CreateUserOutputBoundary)
}

// CreateUserInputData represents the data required for creating a User
type CreateUserInputData struct {
	Username string
}

// CreateUserOutputData represents the data returned from creating a User
type CreateUserOutputData struct {
	ID       int
	Username string
}

// CreateUserOutputBoundary represents a presenter who consumes an output
type CreateUserOutputBoundary interface {
	Present(output CreateUserOutputData)
}

// CreateUserInteractor is responsible for creating Users
type CreateUserInteractor struct{}

// Handle the CreateUserInputData by updating the entities
func (interactor CreateUserInteractor) Handle(input CreateUserInputData, presenter CreateUserOutputBoundary) {
	// TODO: Update entities
	output := CreateUserOutputData{ID: 1, Username: input.Username}
	presenter.Present(output)
}
