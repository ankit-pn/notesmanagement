package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v4"
    "time"

    "notesmanagement/models" 
	"notesmanagement/utils" 
)


func UserSignup(c *gin.Context) {
    var newUser models.User
    if err := c.ShouldBindJSON(&newUser); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    
    hashedPassword, err := utils.HashPassword(newUser.Password)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while hashing password"})
        return
    }
    newUser.Password = string(hashedPassword)

    
    if result := models.DB.Create(&newUser); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
}



func UserLogin(c *gin.Context) {
    var loginDetails models.User
    if err := c.ShouldBindJSON(&loginDetails); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    
    var user models.User
    if result := models.DB.Where("username = ?", loginDetails.Username).First(&user); result.Error != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username or pass"})
        return
    }

    if utils.CheckPasswordHash(loginDetails.Password, user.Password) == false {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username or password"})
        return
    }

    
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": user.ID,
        "exp":     time.Now().Add(time.Hour * 72).Unix(),
    })

    
    tokenString, err := token.SignedString([]byte("your_secret_key")) 
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not log in"})
        return
    }

    
    c.JSON(http.StatusOK, gin.H{"token": tokenString, "user_id": user.ID})
}

