package main

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"log"
)

type CatInfo struct {
	client *elasticsearch.Client
}

func (C CatInfo) Nodes() {
	nodes, e := C.client.Cat.Nodes()
	if e!=nil{
		log.Println(e)
		return
	}
	fmt.Println(nodes.String())
}
func (C CatInfo) Health() {
	health, e := C.client.Cat.Health()
	if e!=nil{
		log.Println(e)
		return
	}
	fmt.Println(health.String())
}
func (C CatInfo) Master() {
	master, e := C.client.Cat.Master()
	if e!=nil{
		log.Println(e)
		return
	}
	fmt.Println(master.String())
}
func (C CatInfo) Indices() {
	indices,e := C.client.Cat.Indices()
	if e!=nil{
		log.Println(e)
		return
	}
	fmt.Println(indices.String())

}

func (C CatInfo) Count() {
	count,e := C.client.Cat.Count(func(request *esapi.CatCountRequest) {
		var v bool
		v = true
		request.V = &v
	})
	if e!=nil{
		log.Println(e)
		return
	}
	fmt.Println(count.String())
}

func ExampleCatAPI() {
	cat := CatInfo{client:DefaultElasticSearchAction.client}
	cat.Nodes()
	cat.Health()
	cat.Master()
	cat.Indices()
	cat.Count()
}