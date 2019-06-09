package main

import "github.com/gin-gonic/gin"

func main() {
	eventListPath := "/event-list"
	r := gin.Default()
	r.GET(eventListPath, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.DELETE(eventListPath, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.PATCH(eventListPath, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST(eventListPath, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
