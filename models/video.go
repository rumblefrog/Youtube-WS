package models

type VideoRequest struct {
	VCode string `form:"vcode" json:"vcode"`
	ITag  int    `form:"itag" json:"itag"`
}
