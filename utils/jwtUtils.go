package utils

import (
    "time"
    "github.com/golang-jwt/jwt/v4"
    "notesmanagement/models" // Replace with your project's package name
)

var jwtKey = []byte("your_secret_key") // Replace with your secret key

// Claims struct will be used to add custom claims in the JWT token
type Claims struct {
    UserID uint
    jwt.StandardClaims
}

// GenerateToken generates a new JWT token for a user.
func GenerateToken(user models.User) (string, error) {
    expirationTime := time.Now().Add(72 * time.Hour)

    // Create the JWT claims, which includes the user ID and expiry time
    claims := &Claims{
        UserID: user.ID,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    // Declare the token with the algorithm used for signing, and the claims
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    // Create the JWT string
    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}

// ValidateToken validates the JWT token and returns the user ID.
func ValidateToken(tokenString string) (*Claims, error) {
    claims := &Claims{}

    // Parse the JWT string and store the result in `claims`
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })

    if err != nil {
        return nil, err
    }

    if !token.Valid {
        return nil, jwt.NewValidationError("invalid token", jwt.ValidationErrorSignatureInvalid)
    }

    return claims, nil
}
