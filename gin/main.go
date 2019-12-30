package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// about get
	r.GET("/about", func(c *gin.Context) {
		admin := c.Query("admin")
		fmt.Println("admin:", admin)
		test := c.Query("test")

		if test == "" {
			test = "admin"
		}
		c.JSON(http.StatusOK, gin.H{"result": 1, "error": "success", "data": gin.H{"admin": admin, "test": test}})
	})

	// http://127.0.0.1:8082/list?list=[11,22,33]
	r.GET("/list", func(c *gin.Context) {
		list := c.QueryArray("list")
		list2 := c.Query("list")
		fmt.Println(list[0][0])
		fmt.Println(list2)
		for _, v := range list {
			fmt.Println(v)
		}
		c.JSON(http.StatusOK, gin.H{"list": list})
	})

	// 获取id
	// http://127.0.0.1:8082/user/349639/admin?info=aaaaa
	r.GET("/user/:id/:name", func(c *gin.Context) {
		id := c.Param("id")
		name := c.Param("name")
		info := c.DefaultQuery("info", "hello world")
		hello := c.DefaultQuery("hello", "hello world")
		fmt.Println("get default info:", info)
		if info == "" {
			info = "hello!"
		}
		c.JSON(http.StatusOK, gin.H{"result": 1, "error": info, "data": gin.H{"id": id, "name": name, "hello": hello}})
	})

	// curl -X POST -d 'admin=yanxxit'  http://127.0.0.1:8082/user/create
	r.POST("/user/create", func(c *gin.Context) {
		admin := c.PostForm("admin")
		c.JSON(http.StatusOK, gin.H{"result": 1, "data": admin})
	})
	r.POST("/user/add", func(c *gin.Context) {
		c.Header("Content-Type", "application/x-www-form-urlencoded")
		admin := c.PostForm("admin")
		c.JSON(http.StatusOK, gin.H{"result": 1, "data": admin})
	})

	// r.Run("http://127.0.0.1:8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	r.Run(":8083") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
