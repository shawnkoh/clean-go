package user

// createUserPresenter is responsible for formatting all OutputData into a ViewModel to be consumed
// by the package's clients.
type createUserPresenter struct {
	viewModel CreateUserViewModel
}

// Present translates the CreateUserOutputData into a ViewModel
func (presenter *createUserPresenter) Present(output createUserOutputData) {
	presenter.viewModel = CreateUserViewModel{ID: output.id, Username: output.username}
}
