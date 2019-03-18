package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/rumblefrog/Youtube-WS/controllers"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.Use(cors.Default())

	ex, err := os.Executable()

	if err != nil {
		panic(err)
	}

	exPath := filepath.Dir(ex)

	router.Use(static.Serve("/", static.LocalFile(exPath+"/dist", false)))

	api := router.Group("/api")
	{
		api.GET("/search", controllers.Search)
		api.GET("/manifest", controllers.Manifest)
		api.GET("/video", controllers.Video)
	}

	log.Println("Started YWS server")

	router.Run(":8080")
}
