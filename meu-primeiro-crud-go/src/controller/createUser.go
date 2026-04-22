package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/matheusdcunha/aprendendo-go/meu-primeiro-crud-go/src/config/logger"
	"github.com/matheusdcunha/aprendendo-go/meu-primeiro-crud-go/src/config/validation"
	"github.com/matheusdcunha/aprendendo-go/meu-primeiro-crud-go/src/controller/model/request"
	"github.com/matheusdcunha/aprendendo-go/meu-primeiro-crud-go/src/model"
)

var (
	UserDomainInterface model.UserDomainInterface
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

	domain := model.NewUserDomain(userRequest.Email, userRequest.Password, userRequest.Name, userRequest.Age)

	if err := domain.CreateUser(); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("user created sucessfully", logger.CreateUserJourney)

	c.String(201, "")
}
