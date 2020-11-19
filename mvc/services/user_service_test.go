package services

import (
	"github.com/azeezkhan2197/golang-microservices/mvc/domain"
	"github.com/azeezkhan2197/golang-microservices/mvc/utils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var (
	userDaoMock usersDaoMock
	getUser     func(userId int64) (*domain.User, *utils.ApplicationError)
)

type usersDaoMock struct{}

func init() {
	domain.UserDao = &usersDaoMock{}
}

func (u *usersDaoMock) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return getUser(userId)
}

func TestGetUserNotFound(t *testing.T) {
	getUser = func(userId int64) (*domain.User, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			StatusCode: http.StatusNotFound,
		}
	}
	user, err := UserService.GetUser(0)
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
}

func TestGetUserFound(t *testing.T) {
	getUser = func(userId int64) (*domain.User, *utils.ApplicationError) {
		return &domain.User{
			Id: 123,
		}, nil
	}
	user, err := UserService.GetUser(0)
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
}
