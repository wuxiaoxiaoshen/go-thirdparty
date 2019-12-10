package main

import (
	"github.com/elastic/go-elasticsearch/v7"
	"log"
)

type ElasticSearchAction struct {
	client *elasticsearch.Client
}

var DefaultElasticSearchAction = NewElasticSearchAction([]string{
	"http://es_01:9200","http://es_02:9200","http://es_03:9200",
})
func NewElasticSearchAction(addresses []string) *ElasticSearchAction {
	return &ElasticSearchAction{client:newEsClient(addresses)}
}
func newEsClient(addresses []string) *elasticsearch.Client {
	cfg := elasticsearch.Config{Addresses: addresses}
	client, err := elasticsearch.NewClient(cfg)

	if err != nil {
		log.Println(err)
		panic(err)
	}

	return client
}