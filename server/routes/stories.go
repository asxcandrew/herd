package routes

import (
	H "github.com/asxcandrew/herd/server/handlers"
	"github.com/gin-gonic/gin"
)

//StoriesRoutes users path group
func StoriesRoutes(route *gin.RouterGroup) {
	story := route.Group("/stories")
	story.POST(":id", H.CreateStory)
}
