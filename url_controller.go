package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin/binding"
	"net/http"
)

type URLController struct {
	//redis connection
	//codec
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
		c.String(http.StatusOK, json.Url + ":" + this.url.Save("http://example.com")) //TODO req["url"]
	}
	// else { invalid request }
}

func (this *URLController) MultiEncode(c *gin.Context) {
	c.String(http.StatusOK, "multi encode")
}

func (this *URLController) Decode(c *gin.Context) {
	c.String(http.StatusOK, "decoded")
}

func (this *URLController) MultiDecode(c *gin.Context) {
	c.String(http.StatusOK, "multi decode")
}