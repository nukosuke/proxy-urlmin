package main

type DecodeRequest struct {
	Url string `form:"url" json:"url" binding:"required"`
}

type MultiDecodeRequest struct {
	Urls []string `form:"urls" json:"urls" binding:"required"`
}
