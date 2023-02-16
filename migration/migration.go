package main

import (
	"github.com/rezanorm/go-rest/initializers"
	"github.com/rezanorm/go-rest/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
