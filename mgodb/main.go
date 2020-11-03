package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type Post struct {
	Name string
	Age int
}

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	collection := client.Database("test").Collection("table1")
	findRecord(ctx, collection)

}

type Collection struct {

}

type Database struct {

}

func insertRecord(ctx context.Context, col *mongo.Collection){
	// 1. 可以构建 bson.D 对象
	// 2. 也可以构建自定义的 struct 对象
	re, e := col.InsertOne(ctx, bson.D{
		{"name", "douyou"},
		{"age", 20},
		{"student", 400},
	})
	if e!=nil{
		log.Println(e)
	}
	fmt.Println(re.InsertedID)

}
func findRecord(ctx context.Context, col *mongo.Collection){
	// 构建 bson.D 对象来进行搜索或者过滤
	var result Post
	filter := bson.D{
		{
			"age", bson.D{
				{"$gt",19},
			},
		},
	}
	// 返回一个
	e := col.FindOne(ctx, filter).Decode(&result)
	if e!=nil{
		log.Println(e)
	}
	fmt.Println("findOne", result)

	filter = bson.D{{
		"age", bson.D{
			{"$in", bson.A{19,20}},
		},
	}}
	// 返回多个，是个游标，可以设置参数
	cursor, e  := col.Find(ctx, filter, options.Find().SetLimit(2))
	if e!=nil{
		log.Println(e)
	}
	for cursor.Next(ctx){
		var result Post
		e = cursor.Decode(&result)
		if e!=nil{
			log.Println(e)
		}
		fmt.Println("findAll", result)
	}

}
func deleteRecord(){}
func updateRecord(){}
