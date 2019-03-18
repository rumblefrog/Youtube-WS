package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rumblefrog/Youtube-WS/models"
	"github.com/rumblefrog/ytdl"
)

func Manifest(c *gin.Context) {
	var m models.ManifestRequest
	var err error

	if err = c.ShouldBindQuery(&m); err != nil {
		c.String(500, "Query binding failed")
		return
	}

	if m.VCode == "" {
		c.String(400, "Missing vcode")
		return
	}

	vi, err := ytdl.GetVideoInfo("https://www.youtube.com/watch?v=" + m.VCode)

	if err != nil {
		c.String(500, "Failed to retrieve video info")
		return
	}

	c.JSON(200, vi)
}
