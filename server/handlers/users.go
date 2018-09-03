package handlers

import (
	"net/http"
	"strconv"

	"github.com/asxcandrew/herd/server/models"
	"github.com/asxcandrew/herd/server/services"
	"github.com/gin-gonic/gin"
)

//GetUser handler for GET users/:username
func GetUser(c *gin.Context) {
	user := &models.User{}
	err := models.DB().Model(user).Where("username = ?", c.Param("username")).Select()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "data": ""})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": user})
}

//UpdateUser handler for PUT users/:username
func UpdateUser(c *gin.Context) {
	var json models.UserResponse
	user := currentUser(c)

	if err := c.BindJSON(&json); err == nil {
		err := services.UpdateUser(user, &json)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "data": err})
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "data": err})
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	if uint64(id) != userID(c) {
		c.AbortWithStatus(http.StatusForbidden)
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
