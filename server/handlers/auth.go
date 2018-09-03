package handlers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/asxcandrew/herd/server/config"
	"github.com/asxcandrew/herd/server/models"
	"github.com/asxcandrew/herd/server/services"
	"github.com/gin-gonic/gin"
)

type signupStruct struct {
	Email    string `json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Username string `form:"username" json:"username" binding:"required"`
	Name     string `form:"name" json:"name"`
}

type loginStruct struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

var authMiddleware *jwt.GinJWTMiddleware

var authResponse = func(c *gin.Context, token string, expire time.Time, user interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"user":   user,
		"code":   http.StatusOK,
		"token":  token,
		"expire": expire.Format(time.RFC3339),
	})
}

//AuthMiddleware provides jwt middleware object
func AuthMiddleware() *jwt.GinJWTMiddleware {
	if authMiddleware == nil {
		authMiddleware = &jwt.GinJWTMiddleware{
			Realm:         "herd",
			Key:           []byte(config.C.AuthSecretKey),
			Timeout:       time.Hour * 24,
			MaxRefresh:    time.Hour,
			Authenticator: Login,
			LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
				c.Request.Header.Set("Authorization", authMiddleware.TokenHeadName+" "+token)

				claims, _ := authMiddleware.GetClaimsFromJWT(c)
				authResponse(c, token, expire, claims["user"])
			},
			PayloadFunc: func(data interface{}) jwt.MapClaims {
				if v, ok := data.(*models.User); ok {
					fmt.Println("PayloadFunc", data)
					return jwt.MapClaims{
						"user":     data,
						"userRole": v.Role,
						"userID":   v.ID,
					}
				}
				return jwt.MapClaims{}
			},
			TokenLookup: "header:Authorization",
			TimeFunc:    time.Now,
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
	var json signupStruct

	if err := c.BindJSON(&json); err == nil {
		user := &models.User{
			Email:    json.Email,
			Username: json.Username,
			Name:     json.Name,
		}
		err := services.Register(user, json.Password)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		mw := AuthMiddleware()
		token, t, _ := mw.TokenGenerator(user)

		authResponse(c, token, t, user)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

// Login handler
func Login(c *gin.Context) (interface{}, error) {
	var json loginStruct
	if err := c.Bind(&json); err != nil {
		return nil, jwt.ErrMissingLoginValues
	}

	user, err := services.Login(json.Email, json.Password)
	if err != nil {
		return nil, jwt.ErrFailedAuthentication
	}

	return user, nil
}
