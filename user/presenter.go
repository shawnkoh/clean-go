package user

// createUserPresenter is responsible for formatting all OutputData into a ViewModel to be consumed
// by the package's clients.
type createUserPresenter struct {
	viewModel CreateUserViewModel
}

// Present implements the createUserOutputBoundary interface.
// It translates CreateUserOutputData into a CreateUserViewModel
func (presenter *createUserPresenter) Present(output createUserOutputData) {
	presenter.viewModel = CreateUserViewModel{ID: output.id, Username: output.username}
}
