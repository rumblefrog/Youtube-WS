package utils

import (
	"log"
	"net/url"
	"strings"

	"github.com/rumblefrog/Youtube-WS/models"
)

// /results?search_query=
// /results?sp=

func ParsePager(s string, active bool) *models.PagerMeta {
	if len(s) < 21 {
		return nil
	}

	meta := &models.PagerMeta{}

	if active {
		s = s[22:]

		var p []string

		p = strings.Split(s, "&amp;sp=")

		if len(p) < 2 {
			return nil
		}

		meta.Query = p[0]
		meta.SP = p[1]
	} else {
		s = s[9:]

		v, err := url.ParseQuery(s)

		if err != nil {
			return nil
		}

		meta.Query = v.Get("search_query")
		meta.SP = v.Get("sp")

		if meta.SP == "" || meta.Query == "" {
			log.Println(s, v)
		}
	}

	return meta
}
