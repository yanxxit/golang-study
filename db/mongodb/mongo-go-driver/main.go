package main

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"golang-study/db/mongodb/mongo-go-driver/model"
	"log"
	"time"
)

//type AccuseLog struct {
//	ID    objectid.ObjectID "_id,omitempty"
//	Name  string	`bson:"dbname",json:"jsonname"`
//	Phone string
//}

func main() {
	accuse_log := model.TingoDb.Collection("accuse_log")
	user := model.TingoDb.Collection("user")
	ctx := model.Ctx

	userData := user.FindOne(ctx, bson.M{"usertoken": "349639"})
	fmt.Println(userData.DecodeBytes())
	userMap := bson.M{}
	userData.Decode(&userMap)
	fmt.Println("----------",userMap["nickname"])//通过map进行获取字段
	list, err := accuse_log.Find(ctx, bson.M{"type": "golang"})
	if err != nil {
		log.Fatal(err)
	}
	for list.Next(ctx) {
		//返回的是map 数据
		var result bson.M
		err := list.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("---->", result["content"])
	}
	if err := list.Err(); err != nil {
		log.Fatal(err)
	}
	cur := time.Now()
	//时间戳：为int32
	timestamp := cur.UnixNano() / (1000000 * 1000)
	fmt.Println(int32(timestamp))
	//accuse_log.InsertOne(ctx, bson.M{
	//	"content": "禅师与青年abcd", "type": "golang", "usertoken": "349639",
	//	"title": "mongo-go-driver", "status": 0, "created": int32(timestamp),
	//})

}
