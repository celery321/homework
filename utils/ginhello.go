package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
	//"net/http"
)

func Ginhello() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})

	r.Run(":80")

}
