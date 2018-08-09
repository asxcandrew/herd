package handlers

import (
	"net/http"

	"github.com/asxcandrew/herd/server/models"
	"github.com/gin-gonic/gin"
)

//CurrentUser handler for users/me
func CurrentUser(c *gin.Context) {
	user := c.Keys["CurrentUser"].(*models.User)

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": user})
}

//UsernameAvalability handler for availability/:username
func UsernameAvalability(c *gin.Context) {
	username := c.Param("username")
	availability, err := models.IsAvailableUsername(username)

	if err != nil || !availability {
		c.JSON(http.StatusNotFound, gin.H{"available": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"available": true})
}
