package main

import (
	"net/http"

	H "github.com/asxcandrew/herd/server/handlers"
	"github.com/asxcandrew/herd/server/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLFiles("server/templates/index.html")

	r.Static("/assets", "server/assets")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", "")
	})

	r.POST("/sign_up", H.Signup)
	r.POST("/sign_in", H.AuthMiddleware().LoginHandler)

	routes.UsersRoutes(r)
	// authorized := r.Group("/a")

	// authorized.Use(authMiddleware.MiddlewareFunc())
	// {

	// }

	r.Run()
}
