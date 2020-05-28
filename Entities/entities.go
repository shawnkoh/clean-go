package entities

// User represents a user account
type User struct {
	id       int
	Username string
}

// Post represents a user account
type Post struct {
	ID   int
	User User
}
