package app

import (
	"github.com/azeezkhan2197/golang-microservices/mvc/controllers"
	"net/http"
)

func StartApp() {
	http.HandleFunc("/user", controller.GetUser)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
