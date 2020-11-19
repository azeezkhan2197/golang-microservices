package services

import (
	"github.com/azeezkhan2197/golang-microservices/mvc/domain"
	"github.com/azeezkhan2197/golang-microservices/mvc/utils"
)

type userService struct{}

var (
	UserService userService
)

func (u *userService) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.UserDao.GetUser(userId)
}
