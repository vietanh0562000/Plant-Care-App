package controllers

import (
	"log"
	"net/http"
	"plant-care-app/database"
	"plant-care-app/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type RegisterInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

func MigratePasswordsToBcrypt() {
	var users []models.User
	database.DB.Find(&users)

	for _, user := range users {
		password := user.Password
		if len(password) > 4 && password[0] == '$' && password[1] == '2' && password[2] == 'a' {
			continue
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

		if err != nil {
			log.Println("Error hashing password", err)
			continue
		}

		user.Password = string(hashedPassword)
		database.DB.Model(&user).Update("password", string(hashedPassword))
	}
}

func Register(c *gin.Context) {
	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while hashing password"})
		return
	}

	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
	}

	database.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}
