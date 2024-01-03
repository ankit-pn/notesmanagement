package main

import (
	"github.com/gin-gonic/gin"
	"notesmanagement/middleware"
	"notesmanagement/models"
	"notesmanagement/routes"
)

func main() {

	gin.SetMode(gin.ReleaseMode)

	models.ConnectDatabase()

	rateLimiter := middleware.NewRateLimiter()

	router := routes.SetupRouter()

	router.Use(middleware.RateLimitMiddleware(rateLimiter, 100))

	router.Run(":8080")
}
