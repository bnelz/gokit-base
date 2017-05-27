package users

import (
	errs "github.com/Boxx/gokit-base/errors"
)

// Service describes the behavior of a user service e.g. CRUD actions
type Service interface {
	// CreateUser defines a new user and returns its id
	CreateUser(id int, fname string, lname string, color string) (int, error)

	// ReadUser finds a user model by id
	ReadUser(id int) (User, error)

	// UpdateUserColor sets a user's favorite color
	UpdateUserColor(id int, color string) error

	// Users returns all users
	Users() []*User
}

// userService is an implementation of the user service interface
type userService struct {
	// Dependencies go here!
	// In the canonical blog example one dependency could be:
	// likes likes.Service
	// The grants service may return all of the post "likes" this user has.
	// Adding this dependency to our user service would allow us to quickly retrieve these likes and compose
	// our user model with them. We may attach a "TotalLikes" field to our business object in this case.

	// userRepo is our user store
	userRepo Repository
}

// NewService returns a new userService
func NewService(repo Repository) Service {
	return &userService{
		userRepo: repo,
	}
}

// CreateUser validates and sends a message to our user storage with a user to create
func (us *userService) CreateUser(id int, fname string, lname string, color string) (int, error) {
	if id <= 0 {
		return id, errs.ErrInvalidArgument
	}

	u := User{
		ID:            id,
		FirstName:     fname,
		LastName:      lname,
		FavoriteColor: color,
	}

	err := us.userRepo.Store(&u)
	if err != nil {
		return id, err
	}

	return u.ID, nil
}

// ReadUser returns a read-only user model from the underlying user repository
func (us *userService) ReadUser(id int) (User, error) {
	if id <= 0 {
		return User{}, errs.ErrInvalidArgument
	}

	u, err := us.userRepo.Find(id)
	return *u, err
}

// Update a user's favorite color in the storage repository
func (us *userService) UpdateUserColor(id int, color string) error {
	if id <= 0 {
		return errs.ErrInvalidArgument
	}

	u, err := us.userRepo.Find(id)
	if err != nil {
		return err
	}

	u.FavoriteColor = color

	return us.userRepo.Store(u)
}

// Users returns all registered users for the application from the repository
func (us *userService) Users() []*User {
	allUsers := us.userRepo.FindAll()
	// Copy the struct
	return allUsers
}
