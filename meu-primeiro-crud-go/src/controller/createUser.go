package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/matheusdcunha/aprendendo-go/meu-primeiro-crud-go/src/config/resterr"
	"github.com/matheusdcunha/aprendendo-go/meu-primeiro-crud-go/src/controller/model/request"
	"github.com/matheusdcunha/aprendendo-go/meu-primeiro-crud-go/src/controller/model/response"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		resterr := resterr.NewBadRequestError(
			fmt.Sprintf("there are some incorrect fields, error=%s", err.Error()))

		c.JSON(resterr.Code, resterr)
		return
	}

	userResponse := response.UserResponse{
		Id:    "1",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}

	c.JSON(201, userResponse)
}
