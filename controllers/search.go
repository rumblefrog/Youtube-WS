package controllers

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
	"github.com/rumblefrog/Youtube-WS/models"
	"github.com/rumblefrog/Youtube-WS/utils"
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

	searchCollector := colly.NewCollector()

	result := &models.Result{}

	searchCollector.OnHTML("li > div.yt-lockup > div.yt-lockup-dismissable", func(e *colly.HTMLElement) {
		video := &models.Video{}

		video.VCode = strings.Split(e.ChildAttr("div.yt-lockup-thumbnail a.yt-uix-sessionlink", "href"), "?v=")[1]
		video.Thumbnail = e.ChildAttr("div.yt-lockup-thumbnail span.yt-thumb-simple img", "src")

		if video.Thumbnail[0:4] == "/yts" {
			video.Thumbnail = e.ChildAttr("div.yt-lockup-thumbnail span.yt-thumb-simple img", "data-thumb")
		}

		video.Title = e.ChildAttr("h3.yt-lockup-title > a.yt-uix-tile-link", "title")
		video.Creator = e.ChildText("div.yt-lockup-byline > a")
		video.TimeLapsed = e.ChildText("div.yt-lockup-meta li:first-child")
		video.Views = e.ChildText("div.yt-lockup-meta li:nth-child(2)")
		video.Description = e.ChildText("div.yt-lockup-description")

		result.Videos = append(result.Videos, video)
	})

	searchCollector.OnHTML("div.search-pager", func(e *colly.HTMLElement) {
		result.CurrentPage, _ = strconv.ParseInt(e.ChildText("button.yt-uix-button > span.yt-uix-button-content"), 10, 64)

		cp := &models.Pager{}

		cp.Meta = utils.ParsePager(e.ChildAttr("button.yt-uix-button", "data-redirect-url"), true)
		cp.Page, _ = strconv.ParseInt(e.ChildText("button.yt-uix-button > span.yt-uix-button-content"), 10, 64)

		result.Pagination = append(result.Pagination, cp)

		e.ForEach("a.yt-uix-button", func(i int, el *colly.HTMLElement) {
			pager := &models.Pager{}

			pager.Meta = utils.ParsePager(el.Attr("href"), false)
			pager.Page, _ = strconv.ParseInt(el.ChildText("span.yt-uix-button-content"), 10, 64)

			result.Pagination = append(result.Pagination, pager)
		})
	})

	searchCollector.OnScraped(func(r *colly.Response) {
		c.JSON(200, result)
	})

	searchCollector.Visit("https://www.youtube.com/results?search_query=" + q.Keyword)
}
