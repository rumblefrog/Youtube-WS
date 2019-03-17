package main

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/rumblefrog/Youtube-WS/controllers"
)

func main() {
	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("./dist", false)))

	api := router.Group("/api")
	{
		api.GET("/search", controllers.Search)
	}

	router.Run(":8080")
}
