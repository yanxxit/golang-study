package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
	"net/http"
	"time"
)

func IndexApi(c *gin.Context) {
	fmt.Println("IndexApi")
	c.JSON(http.StatusOK, gin.H{"result": 1, "data": "indexapi"})
}

// 如果没有注册就使用MustGet方法读取c的值将会抛错，可以使用Get方法取而代之。
func Ping(c *gin.Context) {
	// 获取中间件中的request
	req1 := c.MustGet("request").(string)
	req, _ := c.Get("request")
	fmt.Println("auth:", req1, req)
	c.JSON(200, gin.H{"message": "pong", "request": req})
}

/** 跳转到baidu :重定向 */
func RedirectBaidu(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "http://wwww.baidu.com")
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

func UserCreate(c *gin.Context) {
	admin := c.PostForm("admin")
	version := c.DefaultPostForm("version", "v1.0.1") // 此方法可以设置默认值
	c.JSON(http.StatusOK, gin.H{"result": 1, "data": admin, "version": version})
}

func UserAdd(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	admin := c.PostForm("admin")
	c.JSON(http.StatusOK, gin.H{"result": 1, "data": admin})
}

func GetUserName(c *gin.Context) {
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

func LoginIn(c *gin.Context) {
	var err error
	var user map[string]interface{}
	contentType := c.Request.Header.Get("Content-Type")
	fmt.Println("contentType:", contentType)
	switch contentType {
	case "application/json":
		err = c.BindJSON(&user)
	case "application/x-www-form-urlencoded":
		err = c.BindWith(&user, binding.Form)

	}
	aa := c.Request.FormValue("admin")
	bb := c.Request.PostFormValue("admin")
	fmt.Println("====>", aa, "bb:", bb)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{"result": 1, "error": "success", "data": user})
}

func Login(c *gin.Context) {
	var user map[string]interface{}
	err := c.Bind(&user)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{"result": 1, "error": "success", "data": user})
}

// 等待时间
func Sleep(c *gin.Context) {
	time.Sleep(5 * time.Second)
	c.JSON(http.StatusOK, gin.H{"sleep": "5s"})
}

// 异步协程
func Async(c *gin.Context) {
	cCp := c.Copy()
	// 在请求的时候，sleep5秒钟，同步的逻辑可以看到，服务的进程睡眠了。异步的逻辑则看到响应返回了，然后程序还在后台的协程处理。
	go func() {
		time.Sleep(5 * time.Second)
		log.Println("Done! in path" + cCp.Request.URL.Path)
	}()
	c.JSON(http.StatusOK, gin.H{"async": "ok"})
}
