package routes

import (
	H "github.com/asxcandrew/herd/server/handlers"
	"github.com/asxcandrew/herd/server/middlewares"
	"github.com/gin-gonic/gin"
)

func ApiRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	public := v1.Group("/p")

	public.POST("/sign_up", H.Signup)
	public.POST("/sign_in", H.AuthMiddleware().LoginHandler)
	public.GET("/username_availability/:username", H.UsernameAvalability)
	public.GET("/feed", H.Feed)

	authorized := v1.Group("/a")

	authorized.Use(H.AuthMiddleware().MiddlewareFunc())
	// authorized.Use(middlewares.CurrentUser())
	authorized.Use(middlewares.Authorizer())
	{
		UsersRoutes(authorized)
		StoriesRoutes(authorized)
	}
}
