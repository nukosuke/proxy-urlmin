package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin/binding"
	"github.com/garyburd/redigo/redis"
	"net/http"
)

type URLController struct {
	url *URL
}

func NewURLController(redis_connection redis.Conn) *URLController {
	this := new(URLController)
	this.url = NewURL(redis_connection)
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
		//res, err := URL.Find(json.Url)
		
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"url":    json.Url,
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
