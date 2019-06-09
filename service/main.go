package main

import (
	"fmt"

	"github.com/eventlist/config"
	"github.com/eventlist/service/eventsvc"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	// Setup Database
	postgresConfig := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		config.DBHost, config.Port,
		config.DBName, config.Username,
		config.Password, "disable")
	db, err := gorm.Open("postgres", postgresConfig)
	if err != nil {
		panic(fmt.Sprintf("Can not create db client: %v", err))
	}
	defer db.Close()

	// Setup Gin Handler
	eventListPath := "/event-list/"
	r := gin.Default()

	// Init Handler
	r.GET(eventListPath, func(c *gin.Context) {
		eventsvc.GetEventListEndpoint(db, c)
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.DELETE(eventListPath, func(c *gin.Context) {
		eventsvc.DeleteEventListEndpoint(db, c)
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.PATCH(eventListPath, func(c *gin.Context) {
		eventsvc.UpdateEventListEndpoint(db, c)
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST(eventListPath, func(c *gin.Context) {
		eventsvc.PostEventListEndpoint(db, c)
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
