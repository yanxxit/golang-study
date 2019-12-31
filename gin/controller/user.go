package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexApi(c *gin.Context) {
	fmt.Println("IndexApi")
	c.JSON(http.StatusOK, gin.H{"result": 1, "data": "indexapi"})
}

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}

func About(c *gin.Context) {
	admin := c.Query("admin")
	fmt.Println("admin:", admin)
	test := c.Query("test")

	if test == "" {
		test = "admin"
	}
	c.JSON(http.StatusOK, gin.H{"result": 1, "error": "success", "data": gin.H{"admin": admin, "test": test}})
}

func List(c *gin.Context) {
	list := c.QueryArray("list")
	list2 := c.Query("list")
	fmt.Println(list[0][0])
	fmt.Println(list2)
	for _, v := range list {
		fmt.Println(v)
	}
	c.JSON(http.StatusOK, gin.H{"list": list})
}

func UserCreate(c *gin.Context)  {
	admin := c.PostForm("admin")
	version := c.DefaultPostForm("version", "v1.0.1") // 此方法可以设置默认值
	c.JSON(http.StatusOK, gin.H{"result": 1, "data": admin, "version": version})
}

func UserAdd(c *gin.Context)  {
	c.Header("Content-Type", "application/json")
	admin := c.PostForm("admin")
	c.JSON(http.StatusOK, gin.H{"result": 1, "data": admin})
}

func GetUserName(c *gin.Context)  {
	id := c.Param("id")
	name := c.Param("name")
	info := c.DefaultQuery("info", "hello world")
	hello := c.DefaultQuery("hello", "hello world")
	fmt.Println("get default info:", info)
	if info == "" {
		info = "hello!"
	}
	c.JSON(http.StatusOK, gin.H{"result": 1, "error": info, "data": gin.H{"id": id, "name": name, "hello": hello}})
}
