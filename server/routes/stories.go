package routes

import (
	H "github.com/asxcandrew/herd/server/handlers"
	"github.com/gin-gonic/gin"
)

//StoriesRoutes users path group
func StoriesRoutes(route *gin.RouterGroup) {
	story := route.Group("/stories")
	story.POST("/", H.CreateStory)
	story.GET("/", H.QueryStories)
	story.GET("/:id", H.GetStory)
	story.GET("/:id/body", H.GetStoryBody)
	story.DELETE("/:id", H.DeleteStory)
}
