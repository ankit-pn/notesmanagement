package tests

import (
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gin-gonic/gin"
    "notesmanagement/controllers"
    "encoding/json"
    "bytes"
)

func TestUserSignup(t *testing.T) {
    // Create a new Gin router
    router := gin.Default()
    router.POST("/api/auth/signup", controllers.UserSignup)

    // Create a test request with JSON data
    reqBody := map[string]string{"Username": "testuser", "Password": "testpassword"}
    reqJSON, _ := json.Marshal(reqBody)

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("POST", "/api/auth/signup", bytes.NewBuffer(reqJSON))
    req.Header.Set("Content-Type", "application/json")
    router.ServeHTTP(w, req)

    // Check the response status code
    if w.Code != http.StatusOK {
        t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
    }

    // You can add more assertions to verify the response content here.
}
