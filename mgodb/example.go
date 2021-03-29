package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type ClientForRemote struct {
	session *mgo.Session
}

func (c *ClientForRemote) Dial() {
	s , e := mgo.Dial("mongodb://localhost:27017")
	if e!=nil{
		log.Println(e)
	}
	c.session = s
}

type Log struct {

}

func (L Log)Output(dep int, s string) error{
	log.SetFlags(log.Lshortfile)
	return log.Output(dep,s)
}

type Result struct {
	Id bson.ObjectId `json:"_id" bson:"_id"`
	UserId int `json:"user_id" bson:"user_id"`
}

func TargetExample() {
	//mgo.SetLogger(new(Log))
	//mgo.SetDebug(true)
	db := "FEWeb"
	table := "users"
	var client ClientForRemote
	client.Dial()
	tb := client.session.DB(db).C(table)
	whereBson := bson.M{
		"user_id": bson.M{"$in": []int{67, 9}},
		//"$or": []bson.M{
		//	bson.M{"user_id":1384},
		//	bson.M{"user_id":1385},
		}
	var results []Result
	e := tb.Find(whereBson).Select(bson.M{"_id":1, "user_id":1}).All(&results)
	fmt.Println(e, len(results))
	for i,j := range results{
		fmt.Println(i, j, j.Id.Hex(), len(j.Id.Hex()), 		len(bson.NewObjectId().Hex()))

	}
	//tb.Insert()
	//tb.Bulk().Insert()
	//tb.Bulk().Run()
	//tb.Insert()
	//whereBson = bson.M{}
	//updateBson := bson.M{"$set": bson.M{"filed":1}}
	//tb.Update(whereBson, updateBson)
	//var one =make([]bson.M,0)
	//one = append(one, bson.M{"filed1":1})
	//one = append(one, bson.M{"filed2":1})
	//one = append(one, bson.M{"filed3":1})
	//one = append(one, bson.M{"filed4":1})
	//one = append(one, bson.M{"filed5":1})
	//find := bson.M{"$and": one}
	// 多条件可以将条件写在同一个 bson.M 中，当做其中的一个字段
	// 多条件也可以使用 $and 操作连接多个 bson.M

}