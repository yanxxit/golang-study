package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang-study/gin/controller"
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

	f,_:=os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	v1 := r.Group("/v1")
	{
		// http://127.0.0.1:8083/v1/login
		v1.GET("/login", controller.Ping)
	}

	r.GET("/ping", controller.Ping)

	// about get
	r.GET("/about", controller.About)

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
