package tests

import (
	"net/http"
	"net/http/httptest"
	"plant-care-app/database"
	"plant-care-app/routes"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetlant(t *testing.T) {
	t.Log("Starting GetPlants test...")

	database.Connect()
	t.Log("Database connected")

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	routes.SetupRoutes(router)
	t.Log("Router setup completed")

	token := GetAuthToken(router)
	t.Logf("Got auth token: %s", token)

	req, _ := http.NewRequest("GET", "/plants", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	t.Log("Request created with auth header")

	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)
	t.Logf("Response status: %d", response.Code)
	t.Logf("Response body: %s", response.Body.String())

	assert.Equal(t, http.StatusOK, response.Code)
	t.Log("Test completed successfully")
}
