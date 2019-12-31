package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang-study/gin/controller"
	"golang-study/gin/middleware"
	"io"
	"os"
)

// https://www.jianshu.com/p/98965b3ff638/
func main() {
	r := gin.Default()

	// 使用Logger中间件
	r.Use(gin.Logger())
	// 使用 Recovery 中间件
	r.Use(gin.Recovery())

	// 中间件
	r.Use(middleware.MiddleWare())

	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	// 使用花括号包含被装饰的路由函数只是一个代码规范，即使没有被包含在内的路由函数，只要使用router进行路由，都等于被装饰了。想要区分权限范围，可以使用组返回的对象注册中间件。
	v1 := r.Group("/v1", middleware.ApiV1())
	v1.Use(middleware.ApiV12())
	{
		// curl -X POST -d '{admin:"node.js"}' http://127.0.0.1:8083/v1/login
		// curl --location --request POST 'http://127.0.0.1:8083/v1/login' --header 'Content-Type: application/json' --data-raw '{"admin": "123123"}'
		// 单个路由中间件 middleware.Auth()
		v1.POST("/login", middleware.Auth(), controller.Login)
		v1.POST("/loginIn", middleware.Auth(), controller.LoginIn)
	}

	r.GET("/ping", controller.Ping)
	r.GET("/sleep", controller.Sleep)
	r.GET("/async", controller.Async)

	// about get
	r.GET("/about", controller.About)
	r.GET("/baidu", controller.RedirectBaidu)

	// http://127.0.0.1:8082/list?list=[11,22,33]
	r.GET("/list", controller.List)

	// http://127.0.0.1:8083/user/349639/admin?info=aaaaa
	r.GET("/user/:id/:name", controller.GetUserName)

	// curl -X POST -d 'admin=yanxxit' http://127.0.0.1:8083/user/create
	r.POST("/user/create", controller.UserCreate)
	r.POST("/user/add", controller.UserAdd)

	// http://127.0.0.1:8083/home/info
	r.GET("/home/info", controller.IndexApi)

	// r.Run("http://127.0.0.1:8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	fmt.Println("http://127.0.0.1:8083")
	r.Run(":8083") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
