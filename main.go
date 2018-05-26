package main

import (
	"net/http"

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

	routes.ApiRoutes(r)

	r.Run()
}
