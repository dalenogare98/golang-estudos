package main

import (
	"crud-go/handlers/rest/controllers"
	"crud-go/initializers"
	"github.com/labstack/echo/v4"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}


func main() {
	r := echo.New()

	r.POST("/posts", controllers.PostsCreate)
	r.PUT("/posts/:id", controllers.PostsUpdate)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostsShow)
	r.DELETE("/posts/:id", controllers.PostsDelete)

	r.POST("/users", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.GET("/validate", controllers.Validate)

	r.Logger.Fatal(r.Start(":3000"))
}
