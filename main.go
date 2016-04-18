package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    
    //TODO:
    // create redis connection
    // create codec instance
    //codec := 

    // check wether server is alive
    router.GET("/ping", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "status": http.StatusOK,
            "message": "pong",
        })
    })
    
    // encode API
    router.POST("/encode", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            
        })
    })

    // decode API
    // ^https?://domain.com/:encoded
    router.POST("/decode", func(c *gin.Context) {
        
    })
    
    router.Run()
}