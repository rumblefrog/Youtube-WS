package models

type ManifestRequest struct {
	VCode string `form:"vcode" json:"vcode"`
}
