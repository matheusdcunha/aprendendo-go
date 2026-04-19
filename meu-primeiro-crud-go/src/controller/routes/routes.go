package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/matheusdcunha/aprendendo-go/meu-primeiro-crud-go/src/controller"
)

func InitRoutes(r *gin.RouterGroup) {

	r.GET("/users/:id", controller.FindUserById)
	r.GET("/users/email/:email", controller.FindUserByEmail)
	r.POST("/users", controller.CreateUser)
	r.PUT("/users/:id", controller.UpdateUserById)
	r.DELETE("/users/:id", controller.DeleteUseById)

}
