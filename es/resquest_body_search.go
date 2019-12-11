package main

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"log"
	"strings"
)

type SearchInfo struct {
	client *elasticsearch.Client
}

func (S SearchInfo) Search(index string) {
	response,e := S.client.Search(
		S.client.Search.WithIndex(index),
		S.client.Search.WithBody(strings.NewReader(`{
		  "query": {
			"match_all": {}
		  },
		  "sort": [
			{
			  "hot.keyword": {
				"order": "desc"
			  }
			}
		  ],
	  "from": 0,
	  "size": 10
	}
`)), S.client.Search.WithPretty())
	if e!=nil{
		log.Println(e)
		return
	}
	fmt.Println(response.String())
}


func ExampleSearchInfo() {
	s := SearchInfo{client:DefaultElasticSearchAction.client}
	s.Search("java")
}