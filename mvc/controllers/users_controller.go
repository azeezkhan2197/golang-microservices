package controller

import (
	"github.com/azeezkhan2197/golang-microservices/mvc/services"
	"github.com/azeezkhan2197/golang-microservices/mvc/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user_id must  be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}

		//it sends the error back to the client
		c.JSON(apiErr.StatusCode, apiErr)

		return
	}

	log.Println("user_id is ", userId)

	user, apiErr := services.UserService.GetUser(int64(userId))
	if apiErr != nil {
		c.JSON(apiErr.StatusCode, apiErr)
		return
	}

	c.JSON(http.StatusOK, user)

}
