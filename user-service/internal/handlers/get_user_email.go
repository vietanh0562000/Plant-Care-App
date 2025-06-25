package handlers

import (
	"fmt"
	"net/http"
	"plant-care-app/user-service/internal/database"
	"plant-care-app/user-service/internal/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUserEmail(c *gin.Context) {
	userIdStr := c.Param("id")
	fmt.Printf("user id : %v\n", userIdStr)
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var user models.User

	err = database.DB.First(&user, userId).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"email": user.Email})
}
