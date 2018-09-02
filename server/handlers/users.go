package handlers

import (
	"net/http"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/asxcandrew/herd/server/models"
	"github.com/gin-gonic/gin"
)

//GetUser handler for users/:username
func GetUser(c *gin.Context) {
	var user *models.User
	err := models.QueryStories(user).Where("username = ?", c.Param("username")).Select()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "data": ""})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": user})
}

//CurrentUser handler for users/me
func CurrentUser(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	userID := uint64(claims["userID"].(float64))
	user := &models.User{ID: userID}
	err := models.DB().Select(user)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "data": ""})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": user})
}

//UsernameAvalability handler for availability/:username
func UsernameAvalability(c *gin.Context) {
	username := c.Param("username")
	availability := models.IsAvailableUsername(username)

	c.JSON(http.StatusOK, gin.H{"available": availability})
}
