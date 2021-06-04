package lib

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func PrintStr(msg string) {
	fmt.Println(msg)
}

func Serve() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hooray! Go, Gin and Docker (probably) are properly configured!",
		})
	})
	r.Run()
}
