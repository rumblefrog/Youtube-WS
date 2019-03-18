package controllers

import (
	"io"
	"net/http"
	"strconv"

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

	format := findITag(vi.Formats, r.ITag)

	endpoint, err := vi.GetDownloadURL(format)

	if err != nil {
		c.String(500, "Unable to retrieve download endpoint")
		return
	}

	resp, err := http.Get(endpoint.String())

	if err != nil {
		c.String(500, "Unable to retrieve response from endpoint")
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		c.String(resp.StatusCode, "Origin status code")
		return
	}

	c.Header("Content-Disposition", "Attachment;filename=\""+vi.Title+"\"."+getExt(format))
	c.Header("Content-Length", strconv.FormatInt(resp.ContentLength, 10))

	_, err = io.Copy(c.Writer, resp.Body)
}

func findITag(list ytdl.FormatList, tag int) ytdl.Format {
	for _, f := range list {
		if f.Itag == tag {
			return f
		}
	}

	return list[0]
}

func getExt(format ytdl.Format) string {
	if format.VideoEncoding == "" {
		return format.AudioEncoding
	}

	return format.Extension
}
