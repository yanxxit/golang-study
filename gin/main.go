package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang-study/gin/controller"
	"golang-study/gin/middleware"
	"io"
	"net/http"
	"os"
	"time"
)

// https://www.jianshu.com/p/98965b3ff638/
func main() {
	r := gin.Default()

	// 禁用控制台颜色，当你将日志写入到文件的时候，你不需要控制台颜色。
	//gin.DisableConsoleColor()
	// 保持开启日志彩色显示:
	gin.ForceConsoleColor()
	// Logger 中间件将写日志到 gin.DefaultWriter ,即使你设置 GIN_MODE=release 。
	// 默认 gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())
	//  Recovery 中间件从任何 panic 恢复，如果出现 panic，它会写一个 500 错误。
	r.Use(gin.Recovery())

	// 自定义日志格式
	r.Use(gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			params.ClientIP,
			params.TimeStamp.Format(time.RFC1123),
			params.Method,
			params.Path,
			params.Request.Proto,
			params.StatusCode,
			params.Latency,
			params.Request.UserAgent(),
			params.ErrorMessage, )
	}))

	// 中间件
	r.Use(middleware.MiddleWare())

	// HTML模板Template
	r.LoadHTMLGlob("templates/*")

	// 静态文件服务
	// http://127.0.0.1:8083/html/
	r.Static("/html", "./html")
	// http://127.0.0.1:8083/static/
	r.StaticFS("/static", http.Dir("html"))
	r.StaticFile("/favicon.ico", "./html/favicon.ico")

	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	// 使用花括号包含被装饰的路由函数只是一个代码规范，即使没有被包含在内的路由函数，只要使用router进行路由，都等于被装饰了。想要区分权限范围，可以使用组返回的对象注册中间件。
	v1 := r.Group("/v1", middleware.ApiV1())
	v1.Use(middleware.ApiV12())
	{
		// curl -X POST -d '{admin:"node.js"}' http://127.0.0.1:8083/v1/login
		// curl --location --request POST 'http://127.0.0.1:8083/v1/login' --header 'Content-Type: application/json' --data-raw '{"admin": "123123"}'
		// curl 'http://127.0.0.1:8083/v1/post'
		// 单个路由中间件 middleware.Auth()
		v1.POST("/login", middleware.Auth(), controller.Login)
		v1.POST("/loginIn", middleware.Auth(), controller.LoginIn)
		v1.GET("/post", middleware.Auth(), controller.PostHttp)
		// curl http://127.0.0.1:8083/v1/user/349639
		v1.GET("/user/:userid", middleware.Auth(), controller.GetUserByUserId)
		// curl http://127.0.0.1:8083/v1/users
		v1.GET("/users", middleware.Auth(), controller.GetUserList)
	}

	tpl := r.Group("/tpl")
	// http://127.0.0.1:8083/tpl/user
	tpl.GET("/user", controller.HomeAbout)

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
