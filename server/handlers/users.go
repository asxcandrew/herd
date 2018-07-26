package handlers

import (
	"net/http"
	"strconv"

	"github.com/asxcandrew/herd/server/models"
	"github.com/gin-gonic/gin"
)

//CurrentUser handler for users/me
func CurrentUser(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Keys["userID"].(string), 10, 32)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "data": ""})
		return
	}
	user := &models.User{ID: userID}
	err = models.DB().Select(user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "data": ""})
		return
	}

	res := models.FullUserStruct{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Name:     user.Name,
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": res})
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
