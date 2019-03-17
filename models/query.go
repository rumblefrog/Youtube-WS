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

type Pager struct {
	Page int64 `json:"page"`

	Meta *PagerMeta `json:"meta"`
}

type PagerMeta struct {
	Query string `json:"query"`
	SP    string `json:"sp"`
}

type Result struct {
	Videos      []*Video `json:"videos"`
	Pagination  []*Pager `json:"pagination"`
	CurrentPage int64    `json:"currentpage"`
}
