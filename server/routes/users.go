package routes

import (
	H "github.com/asxcandrew/herd/server/handlers"
	"github.com/gin-gonic/gin"
)

//UsersRoutes users path group
func UsersRoutes(route *gin.RouterGroup) {
	user := route.Group("/users")
	user.GET("me", H.CurrentUser)
	// user.GET("/:username", H.GetUser)
}
