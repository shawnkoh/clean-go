package entities

// User is a user account
type User struct {
	ID       int
	Username string
}

// Post is a user account
type Post struct {
	ID   int
	User User
}
