package main

import (
	"github.com/crud-go/initializers"
	"github.com/crud-go/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
