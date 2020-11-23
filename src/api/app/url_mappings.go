package app

import (
	"github.com/azeezkhan2197/golang-microservices/src/api/controllers/polo"
	"github.com/azeezkhan2197/golang-microservices/src/api/controllers/repositories"
)

func mapUrls(){
	router.GET("/marco",polo.Polo)
	router.POST("/repositories",repositories.CreateRepo)
}
