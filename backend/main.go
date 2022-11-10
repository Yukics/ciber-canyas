package main

import (
	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//initialises a router with the default functions.
	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello world!")
	})

	router.POST("/login", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello world!")
	})

	router.POST("/logout", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello world!")
	})

	// starts the server at port 8080
	router.Run(":8080")

}
