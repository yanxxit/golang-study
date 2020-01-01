package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/valyala/fasthttp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang-study/gin/model"
	"golang-study/gin/service"
	"log"
	"net/http"
	"strconv"
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
	// 支持内部和外部的重定向
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

// 获取User信息
func GetUserByUserId(c *gin.Context) {
	userid := c.Param("userid")
	fmt.Println(userid)
	id, _ := strconv.Atoi(userid)
	user := model.TingoDb.Collection("user")
	ctx := model.Ctx
	var body map[string]interface{}
	user.FindOne(ctx, bson.M{"userid": id}).Decode(&body)

	c.JSON(http.StatusOK, gin.H{"result": 1, "data": body, "error": "success"})
}

// 查询Users 用户列表
func GetUserList(c *gin.Context) {
	accuse_log := model.TingoDb.Collection("accuse_log")
	ctx := model.Ctx

	version := service.GetVersion("admin")
	fmt.Println("version:", version)

	list, _ := accuse_log.Find(ctx, bson.M{"type": "golang"}, options.Find().SetLimit(2), options.Find().SetSort(bson.M{"created": -1}))
	var logs []map[string]interface{}
	for list.Next(ctx) {
		var result map[string]interface{}

		list.Decode(&result)
		fmt.Println(result)
		logs = append(logs, result)
	}

	for _, v := range logs {
		fmt.Println("=====>>>>>>", v, "type:", v["type"])
	}

	c.JSON(http.StatusOK, gin.H{"result": 1, "data": logs, "error": "success"})
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
	log.Println("done! in path" + c.Request.URL.Path)
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

func PostHttp(c *gin.Context) {
	url := `http://httpbin.org/post?key=123`

	// 填充表单，类似于net/url
	args := &fasthttp.Args{}
	args.Add("name", "test")
	args.Add("age", "18")

	status, resp, err := fasthttp.Post(nil, url, args)
	if err != nil {
		fmt.Println("请求失败:", err.Error())
		return
	}

	if status != fasthttp.StatusOK {
		fmt.Println("请求没有成功:", status)
		return
	}

	fmt.Println(string(resp))
	var data map[string]interface{}
	json.Unmarshal([]byte(resp), &data)
	c.JSON(http.StatusOK, gin.H{"data": data})
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

type student struct {
	Name string
	Age  int8
}

//HTML 渲染
//使用 LoadHTMLGlob () 或 LoadHTMLFiles ()
func HomeAbout(c *gin.Context) {
	stu1 := &student{Name: "node.js", Age: 18}
	stu2 := &student{Name: "java", Age: 20}
	c.HTML(http.StatusOK, "user.tmpl", gin.H{
		"title":  "gin",
		"stuArr": [2]*student{stu1, stu2},
	})
}
