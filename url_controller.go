package main

import (
    "net/url"
    "github.com/gin-gonic/gin"
    _ "github.com/gin-gonic/gin/binding"
)

type URLController struct {
    //redis connection
    //codec
}

func (this *URLController) Encode(c *gin.Context) {
    
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