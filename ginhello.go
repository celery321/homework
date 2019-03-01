package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	//"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})

	r.Run(":80")

}
