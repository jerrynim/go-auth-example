package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jerrynim/leave/controllers"
	"github.com/jerrynim/leave/models"
	"github.com/joho/godotenv"
)

func init() {
  // Log error if .env file does not exist
  if err := godotenv.Load(); err != nil {
      log.Printf("No .env file found")
  }
}

func main() {
  r := gin.Default()

  models.ConnectDataBase()

  r.Use(gin.Logger())

  r.GET("/books", controllers.FindBooks)
  r.POST("/books", controllers.CreateBook)
  r.GET("/books/:id", controllers.FindBook)
  r.PATCH("/books/:id", controllers.UpdateBook)
  r.DELETE("/books/:id")

  auth := r.Group("/auth")
  {
    auth.POST("/signup",controllers.CreateUser)
  }

   // Handle error response when a route is not defined
   r.NoRoute(func(c *gin.Context) {
    // In gin this is how you return a JSON response
    c.JSON(404, gin.H{"message": "Not found"})
  })

    r.Run()
}