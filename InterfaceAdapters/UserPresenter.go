package interfaceadapters

import "../usecases"

// UserPresenter is responsible for presenting a User
type UserPresenter struct{}

// Present translates the CreateUserOutputData into a UserViewModel
func (presenter UserPresenter) Present(output usecases.CreateUserOutputData) {
	_ = UserViewModel{ID: output.ID, Username: output.Username}
}
