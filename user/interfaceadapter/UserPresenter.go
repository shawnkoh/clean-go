package interfaceadapter

import "example/user/usecase"

// UserPresenter is responsible for presenting a User
type UserPresenter struct {
	ViewModel UserViewModel
}

// Present translates the CreateUserOutputData into a UserViewModel
func (presenter UserPresenter) Present(output usecase.CreateUserOutputData) {
	presenter.ViewModel = UserViewModel{ID: output.ID, Username: output.Username}
}
