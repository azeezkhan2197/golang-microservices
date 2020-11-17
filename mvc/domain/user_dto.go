package domain

import (
	"fmt"
	"github.com/azeezkhan2197/golang-microservices/mvc/utils"
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
)

func GetUser(userId int64) (*User, *utils.ApplicationError) {
	user, found := users[userId]
	if !found {
		return nil, &utils.ApplicationError{
			Message:    fmt.Sprintf("user doesnot exist with user id : ", userId),
			StatusCode: http.StatusNotFound,
			Code:       "not found",
		}
	}
	return user, nil
}
