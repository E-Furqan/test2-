package models

import (
	"github.com/gin-gonic/gin"

	"github.com/E-Furqan/wy-web-app/config"
	"github.com/E-Furqan/wy-web-app/controllers"
)

func main() {
	config.ConnectDatabase()

	r := gin.Default()

	r.GET("/users", controllers.GetUsers)
	r.GET("/users/:id", controllers.GetUser)
	r.POST("/users", controllers.CreateUser)
	r.PUT("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)

	r.Run() // Defaults to :8080
}
