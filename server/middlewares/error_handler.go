package middlewares

import (
	"github.com/gin-gonic/gin"
)

const errorHandler = "errorHandler"

// ErrorHandler returns middleware.
func ErrorHandler() gin.HandlerFunc {
	applyToContext := func(c *gin.Context) {
		c.Next()
		//TODO: Enable logger here
	}
	return func(c *gin.Context) {
		applyToContext(c)
	}
}
