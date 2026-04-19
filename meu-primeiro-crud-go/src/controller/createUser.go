package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/matheusdcunha/aprendendo-go/meu-primeiro-crud-go/src/config/resterr"
)

func CreateUser(c *gin.Context) {

	err := resterr.NewBadRequestError("implementation not finished")

	c.JSON(err.Code, err)
}
