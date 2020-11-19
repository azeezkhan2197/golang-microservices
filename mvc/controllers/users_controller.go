package controller

import (
	"encoding/json"
	"github.com/azeezkhan2197/golang-microservices/mvc/services"
	"github.com/azeezkhan2197/golang-microservices/mvc/utils"
	"log"
	"net/http"
	"strconv"
)

func GetUser(response http.ResponseWriter, request *http.Request) {
	userId, err := strconv.ParseInt(request.URL.Query().Get("user_id"), 10, 64)

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user_id must  be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}

		jsonValue, _ := json.Marshal(apiErr)
		response.Write(jsonValue)
		return
	}

	log.Println("user_id is ", userId)

	user, apiErr := services.UserService.GetUser(int64(userId))
	if apiErr != nil {
		jsonValue, _ := json.Marshal(apiErr)
		response.WriteHeader(apiErr.StatusCode)
		response.Write(jsonValue)

		return
	}

	jsonValue, _ := json.Marshal(user)
	response.Write(jsonValue)
}
