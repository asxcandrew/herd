package main

import (
	"time"

	jwt "github.com/appleboy/gin-jwt"
	H "github.com/asxcandrew/herd/server/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	authMiddleware := &jwt.GinJWTMiddleware{
		Realm:         "test zone",
		Key:           []byte("secret key"),
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour,
		Authenticator: H.Auth,
		Authorizator: func(userId string, c *gin.Context) bool {
			if userId == "admin" {
				return true
			}

			return false
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

	r.POST("/sign_up", H.Signup)
	r.POST("/sign_in", authMiddleware.LoginHandler)

	// authorized := r.Group("/a")

	// authorized.Use(authMiddleware.MiddlewareFunc())
	// {

	// }

	r.Run()
}
