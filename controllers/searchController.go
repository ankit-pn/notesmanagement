package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"

    "notesmanagement/models" 
)


func SearchNotes(c *gin.Context) {
    query := c.Query("q")
    userID := c.MustGet("userID").(uint)

    var notes []models.Note
    if result := models.DB.Where("user_id = ? AND (title LIKE ? OR content LIKE ?)", userID, "%"+query+"%", "%"+query+"%").Find(&notes); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": notes})
}
