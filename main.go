package main

import (
	"github.com/gin-gonic/gin"
	"github.com/suchy57/Students/controller"
	"github.com/suchy57/Students/database"
)

func main() {

	database.ConnectDB()
	server := gin.Default()

	server.POST("/register", controller.Register)
	server.POST("/login", controller.Login)
	server.GET("/user", controller.User)
	server.GET("/logout", controller.Logout)

	server.Run()
}
