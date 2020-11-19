package domain

import (
	"fmt"
	"github.com/azeezkhan2197/golang-microservices/mvc/utils"
	"log"
	"net/http"
)

var (
	users = map[int64]*User{
		123: {
			Id:        123,
			FirstName: "Abdul",
			LastName:  "Azeez",
			Email:     "azeezkhan2197@gmail.com",
		},
	}
	UserDao userDaoInterface
)

func init() {
	UserDao = &userDao{}
}

type userDaoInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

type userDao struct{}

func (u *userDao) GetUser(userId int64) (*User, *utils.ApplicationError) {
	log.Println("we are accessing database")
	user, found := users[userId]
	if !found {
		return nil, &utils.ApplicationError{
			Message:    fmt.Sprintf("user doesnot exist with user id : %d", userId),
			StatusCode: http.StatusNotFound,
			Code:       "not_found",
		}
	}
	return user, nil
}
