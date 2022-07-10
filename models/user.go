package models

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	// slice that points to user objects
	users  []*User
	nextID = 1
)
