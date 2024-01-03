package routes

import (
    "github.com/gin-gonic/gin"
    "notesmanagement/controllers" 
    "notesmanagement/middleware"  
)

func SetupRouter() *gin.Engine {
    router := gin.Default()

    
    

    
    auth := router.Group("/api/auth")
    {
        auth.POST("/signup", controllers.UserSignup)
        auth.POST("/login", controllers.UserLogin)
    }

    
    notes := router.Group("/api/notes")
    {
        notes.Use(middleware.AuthMiddleware()) 

        notes.GET("/", controllers.GetNotes)
        notes.GET("/:id", controllers.GetNoteByID)
        notes.POST("/", controllers.CreateNote)
        notes.PUT("/:id", controllers.UpdateNote)
        notes.DELETE("/:id", controllers.DeleteNote)
		notes.POST("/:id/share", controllers.ShareNote)
    }

	

    search := router.Group("/api/search")
    {
        search.Use(middleware.AuthMiddleware()) 
        search.GET("/", controllers.SearchNotes)
    }
    

    return router
}
