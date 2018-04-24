package handlers

import (
	"github.com/asxcandrew/herd/server/services"
	"github.com/gin-gonic/gin"
)

// Signup handler
func Signup(c *gin.Context) {
	// notLoggedIn(c)
	c.JSON(200, gin.H{
		"status": "success",
	})
}

// Auth handler
func Auth(uid string, password string, c *gin.Context) (string, bool) {
	id, err := services.Login(uid, password)

	if err != nil {
		return "", false
	}

	return id, true
}
