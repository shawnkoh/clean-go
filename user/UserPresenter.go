package user

// Presenter implements the CreateUserOutputBoundary interface.
// It is responsible for presenting a User.
type Presenter struct {
	ViewModel ViewModel
}

// Present translates the CreateUserOutputData into a ViewModel
func (presenter Presenter) Present(output CreateUserOutputData) {
	presenter.ViewModel = ViewModel{ID: output.ID, Username: output.Username}
}
