package middlewares

import (
	"net/http"
	"strconv"

	"github.com/asxcandrew/herd/server/models"
	"github.com/gin-gonic/gin"
)

const currentUser = "CurrentUser"

// CurrentUser returns middleware.
func CurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		applyToContext(c)
	}
}

func applyToContext(c *gin.Context) {
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
	c.Set(currentUser, user)
	c.Next()
}
