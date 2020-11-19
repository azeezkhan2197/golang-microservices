package app

import (
	"github.com/azeezkhan2197/golang-microservices/mvc/controllers"
)

func mapUrls() {
	router.GET("/users/:user_id", controller.GetUser)

}
