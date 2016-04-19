package main

type EncodeRequest struct {
	Url string `form:"url" json:"url" binding:"required"`
}

type MultiEncodeRequest struct {
	Urls []string `form:"urls" json:"urls" binding:"required"`
}
