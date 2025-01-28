package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HelloWorld is a function that returns a string containing "hello world".
func helloWorldApi(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "hello world"})
}

// HelloWorld is a function that returns a string containing "hello world".
func HelloWorld() string {
	return "hello world"
}

func main() {
	router := gin.Default()
	router.GET("/", helloWorldApi)

	router.Run("localhost:8080")
}
