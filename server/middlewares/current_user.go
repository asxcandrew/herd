package middlewares

import (
	"net/http"

	"github.com/asxcandrew/herd/server/models"
	"github.com/gin-gonic/gin"
)

const currentUser = "CurrentUser"

// CurrentUser returns middleware.
func CurrentUser() gin.HandlerFunc {
	applyToContext := func(c *gin.Context) {
		if c.Keys["userID"] == nil {
			return
		}
		userID := uint64(c.Keys["userID"].(float64))
		user := &models.User{ID: userID}
		err := models.DB().Select(user)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "data": ""})
			c.Abort()
		}
		c.Set(currentUser, user)
		c.Next()
	}
	return func(c *gin.Context) {
		applyToContext(c)
	}
}
