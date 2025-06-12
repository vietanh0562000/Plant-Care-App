package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"plant-care-app/database"
	"plant-care-app/models"
	"plant-care-app/routes"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	// Initialize database
	database.Connect()

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	routes.SetupRoutes(router)

	body := `{
		"username": "Simon",
		"password": "mattkhaumanh00"
	}`

	req, _ := http.NewRequest("POST", "/login", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	response := httptest.NewRecorder()

	router.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestRegister(t *testing.T) {
	t.Log("Starting register test...")

	database.Connect()
	t.Log("Database connected")

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	routes.SetupRoutes(router)
	t.Log("Router setup completed")

	body := `{
		"name": "Alex",
		"email": "alex012@gmail.com",
		"password": "mkmanh00"
	}`
	t.Logf("Request body: %s", body)

	req, err := http.NewRequest("POST", "/register", strings.NewReader(body))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	t.Log("Request created with headers")

	response := httptest.NewRecorder()
	t.Log("Response recorder created")

	router.ServeHTTP(response, req)
	t.Logf("Response status: %d", response.Code)
	t.Logf("Response body: %s", response.Body.String())

	// Print detailed error information if the test fails
	if response.Code != http.StatusOK {
		t.Logf("Registration failed with status code: %d", response.Code)
		t.Logf("Error response: %s", response.Body.String())
		t.FailNow()
	}

	assert.Equal(t, http.StatusOK, response.Code)

	var user models.User
	result := database.DB.Where("name = ?", "Alex").Delete(&user)
	if result.Error != nil {
		t.Logf("Error cleaning up test user: %v", result.Error)
	} else {
		t.Log("Test user cleaned up successfully")
	}
}

func TestGetPlants(t *testing.T) {
	// Initialize database
	database.Connect()

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	routes.SetupRoutes(router)

	// First login to get token
	loginBody := `{
		"username": "Simon",
		"password": "mattkhaumanh00"
	}`

	loginReq, _ := http.NewRequest("POST", "/login", strings.NewReader(loginBody))
	loginReq.Header.Set("Content-Type", "application/json")
	loginResponse := httptest.NewRecorder()
	router.ServeHTTP(loginResponse, loginReq)

	// Parse login response to get token
	var loginResponseBody map[string]string
	json.Unmarshal(loginResponse.Body.Bytes(), &loginResponseBody)
	token := loginResponseBody["token"]

	// Now test the plants endpoint with the token
	req, _ := http.NewRequest("GET", "/plants", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)

	// Print detailed error information if the test fails
	if response.Code != http.StatusOK {
		t.Logf("Request failed with status code: %d", response.Code)
		t.Logf("Error response: %s", response.Body.String())
		t.FailNow()
	}

	assert.Equal(t, http.StatusOK, response.Code)
}
