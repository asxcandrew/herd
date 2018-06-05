package handlers

import (
	"log"
	"net/http"
	"strconv"
	"time"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/asxcandrew/herd/server/models"
	"github.com/asxcandrew/herd/server/services"
	"github.com/gin-gonic/gin"
)

var authMiddleware *jwt.GinJWTMiddleware

type SignupResponse struct {
	Email    string `json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Username string `form:"username" json:"username" binding:"required"`
}

//AuthMiddleware provides jwt middleware object
func AuthMiddleware() *jwt.GinJWTMiddleware {
	if authMiddleware == nil {
		authMiddleware = &jwt.GinJWTMiddleware{
			Realm:         "herd",
			Key:           []byte("secret key"),
			Timeout:       time.Hour,
			MaxRefresh:    time.Hour,
			Authenticator: Auth,
			Authorizator: func(userId string, c *gin.Context) bool {
				return true
			},
			Unauthorized: func(c *gin.Context, code int, message string) {
				c.JSON(code, gin.H{
					"code":    code,
					"message": message,
				})
			},
			TokenLookup:   "header:Authorization",
			TokenHeadName: "Bearer",
			TimeFunc:      time.Now,
		}
		err := authMiddleware.MiddlewareInit()

		if err != nil {
			log.Fatal(err)
		}
	}
	return authMiddleware
}

// Signup handler
func Signup(c *gin.Context) {
	var json SignupResponse

	if err := c.BindJSON(&json); err == nil {
		user := &models.User{
			Email:    json.Email,
			Username: json.Username,
		}
		err := services.Register(user, json.Password)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		mw := AuthMiddleware()
		userToken := mw.TokenGenerator(strconv.Itoa(int(user.ID)))

		c.JSON(http.StatusOK, gin.H{
			"token":  userToken,
			"expire": mw.TimeFunc().Add(mw.Timeout).Format(time.RFC3339),
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

// Auth handler
func Auth(email string, password string, c *gin.Context) (string, bool) {
	user, err := services.Login(email, password)

	if err != nil {
		return "", false
	}

	return strconv.Itoa(int(user.ID)), true
}
