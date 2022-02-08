package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", getHandler)

	if err := r.Run(":5000"); err != nil {
		log.Fatal(err.Error())
	}
}

func getHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message":  true,
		"message2": "hello world",
	})
}
