package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rumblefrog/Youtube-WS/crawler"
	"github.com/rumblefrog/Youtube-WS/models"
)

func Search(c *gin.Context) {
	var q models.Query
	var err error

	if err = c.ShouldBindQuery(&q); err != nil {
		c.String(500, "Query binding failed")

		return
	}

	if q.Keyword == "" {
		c.String(400, "Missing keyword")
		return
	}

	complete := make(chan *models.Result)

	go crawler.Crawl(&q, complete)

	result := <-complete

	c.JSON(200, result)
}
