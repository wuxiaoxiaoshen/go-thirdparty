package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
)
type ElasticClientAction struct {
	client *elastic.Client
	ctx context.Context
}


var DefaultElasticClientAction = NewElasticClientAction("http://172.26.0.3:9200")

func NewElasticClientAction(address string) *ElasticClientAction {
	ctx := context.TODO()
	client,e := elastic.NewClient(
		elastic.SetURL("http://172.26.0.3:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false),
		)
	if e!=nil{
		log.Println(fmt.Sprintf("elasticSearch client fail: %s",e.Error()))
		return nil
	}
	//info, code, e := client.Ping(address).Do(ctx)
	//if e!=nil{
	//	log.Println("elasticSearch ping fail")
	//	return nil
	//}
	//log.Println(fmt.Sprintf("version: %s, code: %d", info.Version.Number,code))
	return &ElasticClientAction{client: client,ctx:ctx}
}

func (E ElasticClientAction) IndexExists(index string) bool {
	ok, e := E.client.IndexExists(index).Do(E.ctx)
	if e!=nil{
		log.Println(fmt.Sprintf("index: %s not exists", index,))
	}
	return ok
}

func (E ElasticClientAction) GetDoc(index string, id string) interface{} {
	result, e := E.client.Get().Index(index).Id(id).Do(E.ctx)
	if e!=nil{
		log.Println(fmt.Sprintf("elasticSearch index: %s, id: %s not exsits",index, id))
		return nil
	}
	return result.Source
}