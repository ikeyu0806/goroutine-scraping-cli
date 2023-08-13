package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "/app/frontend/index.html", nil)
	})

	r.Run(":80")
}
