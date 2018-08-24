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
	story.GET("/:uid", H.GetStory)
	story.PUT("/:uid", H.UpdateStory)
	story.GET("/:uid/body", H.GetStoryBody)
	story.DELETE("/:uid", H.DeleteStory)
}
