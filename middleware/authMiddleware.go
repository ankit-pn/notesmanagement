package middleware

import (
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v4"

    "notesmanagement/models" 
)


func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
            c.Abort()
            return
        }

        
        tokenString := strings.TrimPrefix(authHeader, "Bearer ")
        if tokenString == authHeader {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Could not find bearer token in Authorization header"})
            c.Abort()
            return
        }

        
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, jwt.NewValidationError("invalid signing method", jwt.ValidationErrorSignatureInvalid)
            }
            return []byte("your_secret_key"), nil 
        })

        
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token: " + err.Error()})
            c.Abort()
            return
        }

        
        if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
            userID := uint(claims["user_id"].(float64))
            var user models.User
            if result := models.DB.First(&user, userID); result.Error != nil {
                c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
                c.Abort()
                return
            }

            
            c.Set("userID", userID)
        } else {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        c.Next()
    }
}
