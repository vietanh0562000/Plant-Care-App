package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetFireBaseToken(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Token": "asdfsdfasdf"})
}
