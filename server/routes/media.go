package routes

import (
	H "github.com/asxcandrew/herd/server/handlers"
	"github.com/gin-gonic/gin"
)

//MediaRoutes users path group
func MediaRoutes(route *gin.RouterGroup) {
	media := route.Group("/media")
	media.POST("/", H.CreateMedia)
	media.GET("/:id", H.GetMedia)
	media.DELETE("/:id", H.DeleteMedia)
}
