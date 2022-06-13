package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/assets", "./assets")
	router.SetTrustedProxies(nil)

	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		name := c.DefaultQuery("name", "Anonymous")
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": fmt.Sprintf("Hello %s!", name),
		})
	})

	router.Run(":8080")
}
