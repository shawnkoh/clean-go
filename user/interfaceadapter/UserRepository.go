package interfaceadapter

import "example/entities"

type userRepository struct{}

func (repository userRepository) createUser(username string) entities.User {
}
