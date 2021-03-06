package services

import (
	"github.com/azeezkhan2197/golang-microservices/mvc/domain"
	"github.com/azeezkhan2197/golang-microservices/mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
