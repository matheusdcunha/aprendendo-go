package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/matheusdcunha/aprendendo-go/meu-primeiro-crud-go/src/config/logger"
	"github.com/matheusdcunha/aprendendo-go/meu-primeiro-crud-go/src/config/validation"
	"github.com/matheusdcunha/aprendendo-go/meu-primeiro-crud-go/src/controller/model/request"
	"github.com/matheusdcunha/aprendendo-go/meu-primeiro-crud-go/src/controller/model/response"
)

func CreateUser(c *gin.Context) {
	logger.Info("init CreateUser controller", logger.CreateUserJourney)
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		resterr := validation.ValidateUserError(err)
		logger.Error("error trying validate user info", err, logger.CreateUserJourney)

		c.JSON(resterr.Code, resterr)
		return
	}

	userResponse := response.UserResponse{
		Id:    "1",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}

	logger.Info("user created sucessfully", logger.CreateUserJourney)

	c.JSON(201, userResponse)
}
