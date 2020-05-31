package user

import "example/entity"

// Repository controls the persistence of Users
type Repository interface {
	CreateUser(username string) *entity.User
}
