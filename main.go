package main

import (
	"github.com/abhishekgautam2808/sampleservice2/controllers"
	"github.com/abhishekgautam2808/sampleservice2/models"
	"github.com/gin-gonic/gin"
)

func main() {

	server := SetupRouter()
	server.Run()
}

func SetupRouter() *gin.Engine {
	models.ConnectDatabase()

	server := gin.Default()

	server.GET("/user", controllers.FindAll)
	server.GET("/user/:id", controllers.GetSingleUser)

	return server
}
