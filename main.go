// main.go
package main

import (
	"github.com/E-Furqan/test2-/config"
	"github.com/E-Furqan/test2-/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	server := gin.Default()

	server.GET("/users", controllers.GetUsers)
	server.GET("/users/:id", controllers.GetUser)
	server.POST("/users", controllers.CreateUser)
	server.PUT("/users/:id", controllers.UpdateUser)
	server.DELETE("/users/:id", controllers.DeleteUser)

	server.Run(":8082") // Defaults to :8080
}
