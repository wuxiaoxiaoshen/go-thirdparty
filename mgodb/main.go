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
	//client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	client2, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27018,localhost:27019,localhost:27020/?replicaSet=rs0"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//err = client.Connect(ctx)
	err = client2.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	//defer client.Disconnect(ctx)
	defer client2.Disconnect(ctx)
	//fmt.Println("database")
	//listDataBaseName(ctx, client)
	//fmt.Println("collections")
	//listCollections(ctx, client)


	//collection := client.Database("test").Collection("table1")
	//findRecord(ctx, collection)
	//deleteRecord(ctx,collection)
	//updateRecord(ctx,collection)
	//aggregate(ctx, collection)
	replicate(ctx, client2)

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

func aggregate(ctx context.Context, col *mongo.Collection){
	pipeline := bson.D{

		{"$group", bson.D{
				{"_id", "$name"},
				{"age", bson.D{
					{
						"$sum", "$age", // $sum, $avg, $max, $min, $first, $last
					},
				}},
		}},
	}
	results, e := col.Aggregate(ctx, mongo.Pipeline{pipeline})
	if e!=nil{
		log.Println(e)
	}
	for results.Next(ctx){
		var a bson.D
		results.Decode(&a)
		fmt.Println(a)
	}
}

func replicate(ctx context.Context, client *mongo.Client){
	/*
	如何构建副本机制？
	1. 启动多个实例
	2. 执行 rs.initiate() 命令
	3. 验证 rs.status(), rs.slaveOK(), rs.secondaryOK()
	*/
	col := client.Database("test").Collection("name")
	var result interface{}
	filter := bson.D{
		{"age",20},
	}
	err := col.FindOne(ctx, filter).Decode(&result)
	if err!=nil{
		log.Println(err)
	}
	fmt.Println("repl", result)

	cursor, e := col.Find(ctx, bson.D{})
	if e!=nil{
		log.Println(e)
	}
	var all []interface{}
	cursor.All(ctx, &all)
	for index, i := range all{
		fmt.Println(index, i)
	}

}