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
	fmt.Println("database")
	listDataBaseName(ctx, client)
	fmt.Println("collections")
	listCollections(ctx, client)


	collection := client.Database("test").Collection("table1")
	findRecord(ctx, collection)
	deleteRecord(ctx,collection)
	updateRecord(ctx,collection)

}



type Database struct {
	Name string
	Size int64
}

func listDataBaseName(ctx context.Context, client *mongo.Client){
	results, e := client.ListDatabases(ctx, bson.D{})
	if e!=nil{
		log.Println(e)
	}
	for index, i :=range results.Databases{
		var one Database
		one.Name = i.Name
		one.Size = i.SizeOnDisk
		fmt.Println(index, one)


	}

}
type Collection struct {
	Name string
}

func listCollections(ctx context.Context, client *mongo.Client){
	// 先获取 database
	// 在获取 collection
	results , _ := client.ListDatabaseNames(ctx, bson.D{})
	for _,i :=range results{
		collections ,_:= client.Database(i).ListCollectionNames(ctx, bson.D{})
		fmt.Println(fmt.Sprintf("db: %s, collections: %v",i,collections))
	}
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

	filter = bson.D{{
		"age", bson.D{
			{"$lte", 19},
		},
	}}
	cursor, e = col.Find(ctx, filter)
	if e!=nil{
		log.Println(e)
	}
	var resultsAll []Post
	e = cursor.All(ctx, &resultsAll)
	if e!=nil{
		log.Println(e)
	}
	fmt.Println(resultsAll)


	filter = bson.D{
		{"$or", bson.A{
			bson.D{{"name","xiewei"}},
			bson.D{{"name", "paul"}},
		}},
	}
	var re []interface{}
	cursor, e = col.Find(ctx, filter)
	if e!=nil{
		log.Println(e)
	}
	cursor.All(ctx, &re)
	for index, i :=range re{
		fmt.Println(index, i)
	}
}
func deleteRecord(ctx context.Context, col *mongo.Collection){
	filter := bson.D{
		{
			"students",1,
		},
	}
	results, e := col.DeleteOne(ctx, filter)
	if e!=nil{
		log.Println(e)
	}
	if results.DeletedCount==0{
		log.Println("record not exists")
	}else{
		log.Println("delete record success")
	}
}
func updateRecord(ctx context.Context, col *mongo.Collection){
	filter := bson.D{
		{
			"name", 1,
		},
	}
	updates := bson.D{{
		"$set", bson.D{{
			"name", "paul",
		}},
	}}
	results, e := col.UpdateOne(ctx, filter, updates)
	if e!=nil{
		log.Println(e)
	}
	if results.MatchedCount==0{
		log.Println("record not found")
		return
	}
	if results.ModifiedCount==0{
		log.Println("records updates fail")
	}else{
		log.Println("record updates success")
	}
}
