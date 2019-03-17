package models

type Query struct {
	Keyword string `form:"keyword"`
	SP      string `form:"sp"`
}

type Video struct {
	Title       string `json:"title"`
	Creator     string `json:"creator"`
	VCode       string `json:"vcode"`
	TimeLapsed  string `json:"timelapsed"`
	Views       string `json:"views"`
	Description string `json:"description"`
	Thumbnail   string `json:"thumbnail"`
}

type Result struct {
	Videos []*Video
}
