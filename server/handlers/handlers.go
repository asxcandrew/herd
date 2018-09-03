package handlers

import (
	jwt "github.com/appleboy/gin-jwt"
	"github.com/asxcandrew/herd/server/models"
	"github.com/gin-gonic/gin"
)

var _currentUser *models.User

func currentUser(c *gin.Context) *models.User {
	if _currentUser == nil {
		_currentUser = &models.User{ID: userID(c)}
		models.DB().Select(_currentUser)
	}

	return _currentUser
}

func userID(c *gin.Context) uint64 {
	claims := jwt.ExtractClaims(c)
	return uint64(claims["userID"].(float64))
}
