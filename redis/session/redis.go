package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/shawflying/beego-common-utils/utils/comutil"
	"time"
)

var namespace = "telecom:"

// 创建 redis 客户端
func createClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       2,
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

func main() {
	client := createClient()
	fmt.Println(client)
	// 这里设置过期时间.
	cookie := "15806111230"
	value := make(map[string]interface{})
	value["name"] = "admin"
	//value["age"] = 19
	//value["address"] = "shanghai"
	//temp := make(map[string]interface{})
	//temp["love"] = "ss"
	//temp["city"] = "nj"
	//value["data"] = temp

	str := comutil.TransInterfaceToString(value)

	err := client.Set(namespace+"session:"+cookie, str, 10000*time.Second).Err()
	fmt.Println(err)
	val, err := client.Get(namespace + "session:" + cookie).Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("返回JSON", cookie, val) // age 的值为21

	var s map[string]interface{}
	json.Unmarshal([]byte(val), &s)
	fmt.Println(comutil.TransInterfaceToString(s))
	mykey := "name"
	fmt.Println("直接获取值：", s[mykey])
	for k, v := range s {
		fmt.Println(k, v)
	}

}
