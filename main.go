package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/lordofthemind/ginJwtAuthPostgres/controllers"
	"github.com/lordofthemind/ginJwtAuthPostgres/initializers"
	"github.com/lordofthemind/ginJwtAuthPostgres/middleware"
)

func init() {
	fmt.Println("Initializing...")
	initializers.LoadEnvVariables()
	fmt.Println("Loaded environment variables")
	initializers.ConnectToDB()
	fmt.Println("Connected to database")
	initializers.SyncDatabase()
	fmt.Println("Synchronized database")
}

func main() {
	fmt.Println("Hello, World!")

	router := gin.Default()
	router.POST("/signup", controllers.Signup)
	router.POST("/login", controllers.Login)
	router.GET("/validate", middleware.RequireAuth, controllers.Validate)

	err := router.Run("localhost:9090")
	if err != nil {
		log.Fatal(err)
	}
}
