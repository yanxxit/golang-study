package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

//type AccuseLog struct {
//	ID    objectid.ObjectID "_id,omitempty"
//	Name  string	`bson:"dbname",json:"jsonname"`
//	Phone string
//}

func main() {
	fmt.Println("连接mongodb")
	mongoose := "mongodb://127.0.0.1:27017"
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	client, err := mongo.Connect(ctx,options.Client().ApplyURI(mongoose))
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database("tingodb")
	accuse_log := db.Collection("accuse_log")

	list, err := accuse_log.Find(ctx, bson.M{"type":"golang"})
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
	accuse_log.InsertOne(ctx,bson.M{
		"content":"禅师与青年","type":"golang","usertoken":"349639",
		"title":"mongo-go-driver","status":0,"created":int32(timestamp),
	})

}
