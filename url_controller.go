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
	var json EncodeRequest
	if c.BindJSON(&json) == nil {
		res, err := this.url.Save(json.Url)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"status": http.StatusOK,
				"url":    res,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
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
		for _, value := range json.Urls {
			res, err := this.url.Save(value)
			if err != nil {
				goto ERROR
			}

			value = res
		}

		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"urls":   json.Urls,
		})
	} else {
		goto ERROR
	}

ERROR:
	c.JSON(http.StatusBadRequest, gin.H{
		"status":  http.StatusBadRequest,
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
			"urls":   json.Urls,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Bad Request",
		})
	}
}
