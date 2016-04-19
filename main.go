package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	//TODO:
	// create redis connection

	// NewURLController(redis)
	var ctrl = NewURLController()

	// ping API
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "pong",
		})
	})

	// encode API
	router.GET("/encode", ctrl.Encode)
	router.POST("/encode", ctrl.Encode)
	router.POST("/encode/multi", ctrl.MultiEncode)

	// decode API
	router.POST("/decode", ctrl.Decode)
	router.POST("/decode/multi", ctrl.MultiDecode)

	router.Run()
}
