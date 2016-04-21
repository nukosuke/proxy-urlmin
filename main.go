package main

import (
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()
	router.LoadHTMLFiles("index.html")

	// create redis connection
	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// create controllers
	var ctrl = NewURLController(conn)

	// index page
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "proxy-urlmin",
		})
	})

	// ping API
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "pong",
		})
	})

	// encode API
	router.POST("/encode", ctrl.Encode)
	router.POST("/encode/multi", ctrl.MultiEncode)

	// decode API
	router.POST("/decode", ctrl.Decode)
	router.POST("/decode/multi", ctrl.MultiDecode)

	router.Run()
}
