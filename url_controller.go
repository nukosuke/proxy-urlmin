package main

import (
    "net/http"
    "net/url"
    "github.com/gin-gonic/gin"
    _ "github.com/gin-gonic/gin/binding"
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

func (this *URLController) Decode(c *gin.Context) {
    
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