package routes

import (
	"net/http"

	"strconv"

	"example.com/mygolangproj/models"
	"github.com/gin-gonic/gin"
)

func saveUsers(c *gin.Context) {
	// Get all events from the database
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	user.Save()
}

func GetAllUsers(c *gin.Context) {
	// Get all events from the database
	users, err := models.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}
	c.JSON(http.StatusOK, users)
}

func GetUserByID(c *gin.Context) {
	// Get all events from the database
	id := c.Param("id")
	// Convert id to int
	userID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	user, err := models.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
		return
	}
	c.JSON(http.StatusOK, user)
}
