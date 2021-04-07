package main

import (
	"TryGin/controller"
	"TryGin/service"

	"github.com/gin-gonic/gin"
)

var (
	userService    service.UserService       = service.New()
	UserController controller.UserController = controller.New(userService)
)

func main() {
	server := gin.Default()

	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "HELLO WORLD",
		})
	})

	server.GET("/user", func(c *gin.Context) {
		c.JSON(200, UserController.FindAll())
	})

	server.POST("/user", func(c *gin.Context) {
		c.JSON(200, UserController.Save(c))
	})

	server.Run(":8080")
}
