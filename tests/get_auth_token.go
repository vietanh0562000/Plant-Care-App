package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetAuthToken(router *gin.Engine) string {
	body := `{
		"username": "Simon",
		"password": "mattkhaumanh00"
	}`

	req, _ := http.NewRequest("POST", "/login", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	response := httptest.NewRecorder()

	router.ServeHTTP(response, req)

	if http.StatusOK != response.Code {
		fmt.Println("Fail to login----")
	}

	fmt.Println("Login response: " + response.Body.String())

	var responseBody map[string]string
	json.Unmarshal(response.Body.Bytes(), &responseBody)

	token := responseBody["token"]
	return token
}
