package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Enjoy Sucking"})
	})
	router.GET("/up/", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	router.Run(":80")
}
