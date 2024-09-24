package routes

import (
	"net/http"

	// "encoding/json"
	"github.com/gin-gonic/gin"
)

func TestP(c *gin.Context) {
	var data struct {
		Name string `json:"name" binding:"required"`
	}

	// Binding the input to the struct and handling validation errors
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Simulate internal server error
	if data.Name == "error" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Submission successful"})

}
