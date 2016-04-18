package main
import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/gin-gonic/gin/binding"
    "./controllers"
)

func main() {
    router := gin.Default()

    // check wether server is alive
    router.GET("/ping", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "status": http.StatusOK,
            "message": "pong"
        })
    })
    
    // encode API
    router.GET("/encode", func(c *gin.Context) {
        c.JSON(http.StatusOK)
    })
    router.POST("/encode", func(c *gin.Context) {
        
    })

    
    // decode API
    router.GET("/decode", func(c *gin.Context) {
        
    })
    router.POST("/decode", func(c *gin.Context) {
        
    })
    
    router.Run()
}