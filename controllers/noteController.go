package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"

    "notesmanagement/models" 
)


func GetNotes(c *gin.Context) {
    var notes []models.Note
    userID := c.MustGet("userID").(uint) 

    if result := models.DB.Where("user_id = ?", userID).Find(&notes); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": notes})
}


func GetNoteByID(c *gin.Context) {
    var note models.Note
    noteID := c.Param("id")
    userID := c.MustGet("userID").(uint) 

    if result := models.DB.Where("id = ? AND user_id = ?", noteID, userID).First(&note); result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": note})
}


func CreateNote(c *gin.Context) {
    var newNote models.Note
    if err := c.ShouldBindJSON(&newNote); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    userID := c.MustGet("userID").(uint) 
    newNote.UserID = userID

    if result := models.DB.Create(&newNote); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"data": newNote})
}


func UpdateNote(c *gin.Context) {
    noteID := c.Param("id")
    userID := c.MustGet("userID").(uint)

    var existingNote models.Note
    if result := models.DB.Where("id = ? AND user_id = ?", noteID, userID).First(&existingNote); result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
        return
    }

    var updatedNoteData models.Note
    if err := c.ShouldBindJSON(&updatedNoteData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    models.DB.Model(&existingNote).Updates(updatedNoteData)
    c.JSON(http.StatusOK, gin.H{"data": existingNote})
}



func DeleteNote(c *gin.Context) {
    noteID := c.Param("id")
    userID := c.MustGet("userID").(uint)

    var note models.Note
    if result := models.DB.Where("id = ? AND user_id = ?", noteID, userID).First(&note); result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
        return
    }

    models.DB.Delete(&note)
    c.JSON(http.StatusOK, gin.H{"message": "Note deleted successfully"})
}






