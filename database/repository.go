package database

import "example/entity"

// UserRepository implements the UserComponent interface
type UserRepository struct{}

// CreateUser creates a User
func (repository UserRepository) CreateUser(username string) *entity.User {
	// TODO: Actually create the user and retrieve id from the database.
	return &entity.User{ID: "abc", Username: username}
}
