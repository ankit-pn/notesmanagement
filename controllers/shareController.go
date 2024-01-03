package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "notesmanagement/models" 
)



func ShareNote(c *gin.Context) {
    noteID := c.Param("id")
    userID := c.MustGet("userID").(uint)

    var note models.Note
    if result := models.DB.Where("id = ? AND user_id = ?", noteID, userID).First(&note); result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
        return
    }

    var shareRequest struct {
        UserIDs []uint `json:"user_ids"` 
    }
    if err := c.ShouldBindJSON(&shareRequest); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    
    var usersToShareWith []models.User
    if result := models.DB.Where("id IN ?", shareRequest.UserIDs).Find(&usersToShareWith); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving users"})
        return
    }

    note.SharedWithUsers = append(note.SharedWithUsers, usersToShareWith...)

    
    if result := models.DB.Save(&note); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }

    
    c.JSON(http.StatusOK, gin.H{"message": "Note shared successfully", "shared_with_users": usersToShareWith})
}

