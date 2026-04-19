package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/matheusdcunha/aprendendo-go/meu-primeiro-crud-go/src/controller/routes"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(port); err != nil {
		log.Fatal(err)
	}
}
