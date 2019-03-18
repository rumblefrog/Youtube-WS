package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/rumblefrog/Youtube-WS/controllers"
)

func main() {
	router := gin.Default()

	router.Use(cors.Default())
	router.Use(static.Serve("/", static.LocalFile("./dist", false)))

	api := router.Group("/api")
	{
		api.GET("/search", controllers.Search)
		api.GET("/manifest", controllers.Manifest)
		api.GET("/video", controllers.Video)
	}

	router.Run(":8080")
}
