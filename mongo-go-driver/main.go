package main

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang-study/mongo-go-driver/model"
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

	// 查询一个数据：做一个唯一查询
	userData := user.FindOne(ctx, bson.M{"usertoken": "349639"})
	userData1 := user.FindOne(ctx, map[string]string{"usertoken": "349639"})
	// 将string 转换为ObjectId
	id, _ := primitive.ObjectIDFromHex("5cfd288afa9aae5de0c25d4b")
	fmt.Println(id, id.Hex())
	// 通过ObjectId 获取数据
	userData2 := accuse_log.FindOne(ctx, bson.M{"_id": id})

	// fmt.Println(userData.DecodeBytes())
	fmt.Println(userData2.DecodeBytes())
	userMap := bson.M{}
	userMap1 := bson.M{}
	userMap2 := bson.M{}
	userData.Decode(&userMap)
	userData1.Decode(&userMap1)
	userData2.Decode(&userMap2)

	fmt.Println("----------", userMap["nickname"])            //通过map进行获取字段
	fmt.Println("----------", userMap["phonenumber"])         //通过map进行获取字段
	fmt.Println("----------", userMap["_id"])                 //通过map进行获取字段
	fmt.Println("----------:userData1", userMap1["nickname"]) //通过map进行获取字段
	fmt.Println("----------:userData2", userMap2["content"])  //通过map进行获取字段
	// ok
	userData3 := user.FindOne(ctx, bson.M{"_id": userMap["_id"]})
	fmt.Println(userData3.DecodeBytes())
	// 1. accuse_log aggregate 聚合查询
	var pipeline []bson.M
	pipeline = []bson.M{
		bson.M{"$match": bson.M{"usertoken": "349639"}},
		bson.M{"$group": bson.M{"_id": "$type", "count": bson.M{"$sum": 1}}},
	}
	// 聚合查询
	aggregate1, _ := accuse_log.Aggregate(ctx, pipeline)

	for aggregate1.Next(ctx) {
		//返回的是map 数据
		var result bson.M
		aggregate1.Decode(&result)
		fmt.Println("aggregate:_id", result["_id"])
		fmt.Println("aggregate:count", result["count"])
	}

	// 查询列表数据
	list, err := accuse_log.Find(ctx, bson.M{"type": "golang"})
	list2, _ := accuse_log.Find(ctx, bson.M{"type": "golang", "content": primitive.Regex{Pattern: "abcd"}})
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
	for list2.Next(ctx) {
		//返回的是map 数据
		var result bson.M
		err := list2.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("2---->", result["content"])
	}
	if err := list.Err(); err != nil {
		log.Fatal(err)
	}
	cur := time.Now()
	//时间戳：为int32
	timestamp := cur.UnixNano() / (1000000 * 1000)
	fmt.Println(int32(timestamp))
	// 插入数据
	accuse_log.InsertOne(ctx, bson.M{
		"content": "即将删除的数据", "type": "golang", "usertoken": "349639",
		"title": "mongo-go-driver", "status": 0, "created": int32(timestamp),
	})

	// 删除数据
	accuse_log.DeleteOne(ctx, bson.M{"content": "即将删除的数据"})

	// 更新数据
	dd := accuse_log.FindOneAndUpdate(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"title": "mongo-update-修改内容"}})
	fmt.Println(dd.DecodeBytes())
}
