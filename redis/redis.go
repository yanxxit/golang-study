package main

import "fmt"
import (
	"github.com/go-redis/redis"
	"time"
)

// 创建 redis 客户端
func createClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		PoolSize: 10,
	})

	// 通过 cient.Ping() 来检查是否成功连接到了 redis 服务器
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	return client
}

//获取缓存信息
func GetCache() {

}

func AddCatch() string {
	namespace := "telecom:"
	key := "17721021494"
	value := `{"status":"200"}`
	expires := 10000 * time.Second
	fmt.Println(namespace, key)
	client := createClient()
	err := client.Set(namespace+key, value, expires).Err()
	if err != nil {
		return `{"status":"0"}`
	} else {
		return `{"status":"200"}`
	}
}

func main() {
	client := createClient()
	fmt.Println(client)
	AddCatch()
	// 这里设置过期时间.
	err := client.Set("age", "20", 10000*time.Second).Err()
	if err != nil {
		panic(err)
	}

	client.Set("JSON", `{"status":200}`, 10000*time.Second).Err()
	client.Set("apis:getName:data", `{"status":200}`, 10000*time.Second).Err()

	val, err := client.Get("JSON").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("返回JSON", "JSON", val) // age 的值为21
}
