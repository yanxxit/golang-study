package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
	"github.com/shawflying/beego-common-utils/utils/logger"
	"github.com/astaxie/beego"
	"strconv"
)

// 创建 redis 客户端
func CreateRedisClient() *redis.Client {
	DB, _ := strconv.Atoi(beego.AppConfig.String("redis_DB"))
	PoolSize, _ := strconv.Atoi(beego.AppConfig.String("redis_PoolSize"))
	client := redis.NewClient(&redis.Options{
		Addr:     beego.AppConfig.String("redis_Addr"),
		Password: beego.AppConfig.String("redis_Password"),
		DB:       DB,
		PoolSize: PoolSize,
	})

	// 通过 cient.Ping() 来检查是否成功连接到了 redis 服务器
	pong, err := client.Ping().Result()
	if err == nil {
		fmt.Println("链接 " + beego.AppConfig.String("redis_Addr") + " redis 数据库成功")
	} else {
		fmt.Println("链接 "+beego.AppConfig.String("redis_Addr")+" redis 数据库，请尽快处理", pong, err)
	}

	return client
}

var ResdisClient *redis.Client

func init() {
	ResdisClient = CreateRedisClient()
}

//命名空间
var NameSpace = "telecom:"

//获取缓存信息
func GetCache(key string) string {
	val, _ := ResdisClient.Get(NameSpace + key).Result()
	logger.Info("获取缓存：", key, "成功 "+val)
	return val
}

//添加缓存
func AddCatch(key string, value interface{}, expires time.Duration) {
	err := ResdisClient.Set(NameSpace+key, value, expires).Err()
	if err == nil {
		logger.Info("添加缓存：", key, value, "成功")
	} else {
		logger.Info("添加缓存：", key, value, "失败")
	}
}

//清除缓存
func DelCatch(key string) {
	err := ResdisClient.Del(NameSpace + key).Err()
	if err == nil {
		logger.Info("清理缓存：", key, "成功")
	} else {
		logger.Info("清理缓存：", key, "失败")
	}
}
