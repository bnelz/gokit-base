// Users package is a sample business domain object package for application users
package users

// User describes an application user business object
type User struct {
	ID            int    `json:"id"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	FavoriteColor string `json:"fav_color,omitempty"`
}

// New returns a reference to a new user instance
func New(id int, fname string, lname string) *User {
	return &User{
		ID:        id,
		FirstName: fname,
		LastName:  lname,
	}
}

// (User) Repository is the set of behavior a repository, or "store", of users must conform to.
type Repository interface {
	// Store a new user in the repository
	Store(user *User) error

	// Find a user in the repository by ID
	Find(id int) (*User, error)

	// FindAll users in the repository
	FindAll() []*User
}
