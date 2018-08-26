package main

import (
	"github.com/asxcandrew/herd/server/middlewares"
	"github.com/asxcandrew/herd/server/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middlewares.ErrorHandler())
	routes.ApiRoutes(r)

	// TODO: Pass port number in env variable
	r.Run(":8080")
}
