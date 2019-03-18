package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rumblefrog/Youtube-WS/models"
	"github.com/rumblefrog/ytdl"
)

func Video(c *gin.Context) {
	var r models.VideoRequest
	var err error

	if err = c.ShouldBindQuery(&r); err != nil {
		c.String(500, "Query binding failed")
		return
	}

	if r.VCode == "" {
		c.String(400, "Missing vcode")
		return
	}

	if r.ITag == 0 {
		c.String(400, "Missing itag")
		return
	}

	vi, err := ytdl.GetVideoInfo("https://www.youtube.com/watch?v=" + r.VCode)

	if err != nil {
		c.Status(500)
		return
	}

	c.Header("Content-Disposition", "Attachment;filename="+vi.Title)

	vi.Download(vi.Formats[0], c.Writer)
}
