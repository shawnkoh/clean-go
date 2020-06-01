package user

// Controller packages input data from the server and passes this object through the
// InputBoundary to the UseCaseInteractor
// It also uses the DataAccessInterface to bring the data used by those
// Entities into memory from the Database.
// Upon completion, the UseCaseInteractor gathers data from the Entities and constructs the OutputData
// as another plain old object.
// The OutputData is then passed through the OutputBoundary interface to the Presenter.
// Anyone who wants to be a controller of user has to implement all the following functions
type Controller interface {
	CreateUser(username string) CreateUserViewModel
}

type controller struct {
	repository Repository
}

// CreateUser creates a User.
func (controller controller) CreateUser(username string) CreateUserViewModel {
	input := createUserInputData{username: username}
	interactor := NewCreateUserInteractor(controller.repository)
	presenter := createUserPresenter{}
	interactor.Handle(input, presenter)
	return presenter.viewModel
}

// NewController creates a UserController
func NewController(repository Repository) Controller {
	return controller{repository: repository}
}
