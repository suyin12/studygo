package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 待办事项结构体
type Todo struct {
	gorm.Model
	Title  string `json:"title" binding:"required"`
	Status int    `json:"status" binding:"required"`
}

func main() {
	//创建一个默认路由引擎
	r := gin.Default()

	//创建数据表todo
	dsn := "root:root@tcp(127.0.0.1:3306)/go_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))

	// 先不管，关闭不了数据库连接
	// defer db.Close()
	if err != nil {
		fmt.Printf("databaes connect failed%v\n", err.Error())
		return
	}

	//数据迁移
	db.AutoMigrate(&Todo{})

	//获取全部数据
	r.GET("/", func(c *gin.Context) {
		var todo_list []Todo
		fmt.Printf("todo type %T\n", todo_list)
		db.Find(&todo_list)
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
			"data":    todo_list,
		})
	})

	//创建待办事项
	r.POST("/todo", func(c *gin.Context) {
		var todo Todo
		if err := c.ShouldBind(&todo); err == nil {
			fmt.Printf("todo data:%#v\n", todo)
			c.JSON(http.StatusOK, gin.H{
				"title":  todo.Title,
				"status": todo.Status,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		db.Create(&todo)
	})

	//更新待办事项
	r.PUT("/todo/:id", func(c *gin.Context) {
		var todo Todo
		id := c.Param("id")

		fmt.Println(id)
		db.First(&todo, id)
		// 注意：下面为了举例子方便，暂时忽略了错误处理
		b, _ := c.GetRawData() // 从c.Request.Body读取请求数据
		// 反序列化
		_ = json.Unmarshal(b, &todo)
		db.Save(&todo)
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
			"data":    todo,
		})
	})

	//删除待办事项
	r.DELETE("/todo/:id", func(c *gin.Context) {
		var todo Todo
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
			"id":      id,
		})
		db.Delete(&todo, id)
	})

	r.Run(":9090")
}
