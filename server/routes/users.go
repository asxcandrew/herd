package routes

import "github.com/gin-gonic/gin"

//UsersRoutes users path group
func UsersRoutes(route *gin.Engine) {
	user := route.Group("/users")
	user.GET("me")
}
