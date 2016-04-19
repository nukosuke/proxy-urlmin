package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin/binding"
	"net/http"
	"net/url"
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
	c.String(http.StatusOK, this.url.Save("http://example.com")) //TODO req["url"]
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

func (this *URLController) validate(url_string string) bool {
	var isValid bool

	_, err := url.Parse(url_string)

	if err != nil {
		isValid = false
	} else {
		isValid = true
	}

	return isValid
}
