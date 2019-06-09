package model

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var TingoDb *mongo.Database
var Ctx context.Context

func init()  {
	fmt.Println("初始化后")
	mongoose := "mongodb://127.0.0.1:27017/tingodb"
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	client, err := mongo.Connect(ctx,options.Client().ApplyURI(mongoose))
	if err != nil {
		log.Fatal(err)
	}
	TingoDb = client.Database("tingodb")
}
