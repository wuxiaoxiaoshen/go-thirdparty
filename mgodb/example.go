package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type ClientForRemote struct {
	session *mgo.Session
}

func (c *ClientForRemote) Dial() {
	s , e := mgo.Dial("mongodb://hw-sg-mildom-test2.livenono.com:27103")
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

func TargetExample() {
	mgo.SetLogger(new(Log))
	mgo.SetDebug(true)
	db := "HostReport"
	table := "host_wage_rule"
	var client ClientForRemote
	client.Dial()
	tb := client.session.DB(db).C(table)
	whereBson := bson.M{
		//"user_id": bson.M{"$in": []int{1384, 1385}},
		"$or": []bson.M{
			bson.M{"user_id":1384},
			bson.M{"user_id":1385},
		},
	}
	var results []interface{}
	tb.Find(whereBson).Limit(10).Sort("-update_time").All(&results)
	for i,j := range results{
		result, _ := json.MarshalIndent(j, " ", " ")
		fmt.Println(i,string(result))
	}
	tb.Insert()
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