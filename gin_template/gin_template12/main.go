package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//创建一个默认路由引擎
	r := gin.Default()

	//单个文件上传
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("f1")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		log.Println(file.Filename)
		dst := fmt.Sprintf("C:/Users/Administrator/Desktop/%s", file.Filename) //字符串拼接，上传保存到桌面
		// 上传文件到指定的目录
		c.SaveUploadedFile(file, dst)
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("'%s' uploaded!", file.Filename),
		})
	})

	//处理多个文件上传
	r.POST("/uploads", func(c *gin.Context) {
		form, _ := c.MultipartForm()

		files := form.File["f1"]
		log.Println(form)
		for index, file := range files {
			log.Println(file.Filename)
			dst := fmt.Sprintf("C:/Users/Administrator/Desktop/%s_%d", file.Filename, index) //字符串拼接，上传保存到桌面
			// 上传文件到指定的目录
			c.SaveUploadedFile(file, dst)
		}

		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("'%d' uploaded!", len(files)),
		})
	})

	//http 重定向
	r.GET("/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://www.sogo.com/")
	})

	//路由重定向
	r.GET("/r-redirect", func(c *gin.Context) {
		// 指定重定向的URL
		c.Request.URL.Path = "/test2"
		r.HandleContext(c)
	})

	r.GET("/test2", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	//启动
	r.Run(":9090")
}
