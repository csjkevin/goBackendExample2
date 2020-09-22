package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goBackendExample/internal/controller"
	"goBackendExample/internal/repository"
)

func init() {
	repository.Setup()
}

func main() {
	fmt.Println("Go backend example starting...")

	// initialize gin framework
	r := gin.Default()

	// route binding
	r.Any("/", controller.Hello)

	user := r.Group("/users")
	{
		user.GET("", controller.GetAllUser)
		user.GET("/:id", controller.GetUserById)
		user.POST("", controller.CreateUser)
		user.PATCH("/:id", controller.UpdateUser)
	}

	// start gin
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	fmt.Println("Go backend example initialized!")
}
