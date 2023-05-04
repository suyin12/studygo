package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "GET",
		})
	})

	r.POST("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "POST",
		})
	})

	r.PUT("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "PUT",
		})
	})

	r.DELETE("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "DELETE",
		})
	})

	//启动Http服务，默认在0.0.0.0:8080启动服务
	r.Run()
}
