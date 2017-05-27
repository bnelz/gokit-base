package users

import (
	"testing"

	"errors"

	errs "github.com/Boxx/gokit-base/errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserService_CreateUserInvalidArgs(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepository(ctrl)
	us := NewService(mockRepo)
	mockUser := User{
		ID:            0,
		FirstName:     "Bob",
		LastName:      "YourUncle",
		FavoriteColor: "Blue",
	}
	id, err := us.CreateUser(mockUser.ID, mockUser.FirstName, mockUser.LastName, mockUser.FavoriteColor)
	assert.Equal(t, mockUser.ID, id)
	assert.EqualError(t, err, errs.ErrInvalidArgument.Error())
}

func TestUserService_CreateUserFailureOnStore(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepository(ctrl)
	us := NewService(mockRepo)
	mockUser := User{
		ID:            1,
		FirstName:     "Bob",
		LastName:      "YourUncle",
		FavoriteColor: "Blue",
	}
	mockRepo.EXPECT().Store(&mockUser).Return(errors.New("I'm a repository error!"))
	id, err := us.CreateUser(mockUser.ID, mockUser.FirstName, mockUser.LastName, mockUser.FavoriteColor)
	assert.Error(t, err)
	assert.Equal(t, mockUser.ID, id)
}

func TestUserService_CreateUserSuccessfulStore(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepository(ctrl)
	mockUser := User{
		ID:            1,
		FirstName:     "Bob",
		LastName:      "YourUncle",
		FavoriteColor: "Blue",
	}
	us := NewService(mockRepo)
	mockRepo.EXPECT().Store(&mockUser).Return(nil)
	id, err := us.CreateUser(mockUser.ID, mockUser.FirstName, mockUser.LastName, mockUser.FavoriteColor)
	assert.NoError(t, err)
	assert.Equal(t, mockUser.ID, id)
}
