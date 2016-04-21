package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin/binding"
	"net/http"
)

type URLController struct {
	//redis connection
	url *URL
}

func NewURLController() *URLController {
	this := new(URLController)
	this.url = NewURL()
	return this
}

func (this *URLController) Encode(c *gin.Context) {
	//TODO: read & validate request
	var json EncodeRequest
	if c.BindJSON(&json) == nil {
		if this.url.Validate(json.Url) {
			c.JSON(http.StatusOK, gin.H{
				"status": http.StatusOK,
				"url":    this.url.Save(json.Url),
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": http.StatusBadRequest,
				"message": "Invalid URL Schema",
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Bad Request",
		})
	}
}

func (this *URLController) MultiEncode(c *gin.Context) {
	var json MultiEncodeRequest
	if c.BindJSON(&json) == nil {
		//TODO:
		// move validation method to URL model
		for _, value := range json.Urls {
			if !this.url.Validate(value) {
				goto ERROR
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"urls":   json.Urls,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Bad Request",
		})
	}
	
ERROR:
	c.JSON(http.StatusBadRequest, gin.H{
		"status": http.StatusBadRequest,
		"message": "Bad Request",
	})
}

func (this *URLController) Decode(c *gin.Context) {
	var json DecodeRequest
	if c.BindJSON(&json) == nil {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"url":    "http://example.com",
		})
	}
	//this.url.Find()
}

func (this *URLController) MultiDecode(c *gin.Context) {
	var json MultiDecodeRequest
	if c.BindJSON(&json) == nil {
		//TODO:
		//map(json.Urls, func(i) { this.url.Find(i) });
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"urls": json.Urls,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"message": "Bad Request",
		})
	}
}