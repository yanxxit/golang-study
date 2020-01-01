package model

import (
	"context"
	"fmt"
	"golang-study/gin/config"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var TingoDb *mongo.Database
var Ctx context.Context

func init() {
	fmt.Println("初始化后")
	// mongoose := "mongodb://127.0.0.1:27017/tingodb"
	mongoose := config.Conf.Mongodb
	//mongoose := "mongodb://tingo.w:w#tingo@172.16.1.102:27017/tingodb"
	fmt.Println(mongoose)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	// 链接mongo服务
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoose))
	if err != nil {
		log.Fatal(err)
	}
	// 判断服务是否可用
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	// 选择数据库
	TingoDb = client.Database("tingodb")
}
