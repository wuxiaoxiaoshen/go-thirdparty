package main

import (
	"context"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"log"
	"strings"
)

type ElasticSearchAction struct {
	client *elasticsearch.Client
}

var DefaultElasticSearchAction = NewElasticSearchAction([]string{
	"http://0.0.0.0:9200",
})

func NewElasticSearchAction(addresses []string) *ElasticSearchAction {
	return &ElasticSearchAction{client: newEsClient(addresses)}
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

func (E ElasticSearchAction) Version() string {
	return elasticsearch.Version
}

func (E ElasticSearchAction) IndexIsExists(index string) bool {
	response, e := E.client.Indices.Exists([]string{index}, func(request *esapi.IndicesExistsRequest) {
		request.Pretty = true
		request.Human = true
		request.Index = []string{index}
	})
	if e != nil {
		log.Println(e)
		return false
	}
	defer response.Body.Close()
	log.Println(response.String())
	if response.StatusCode == 200 {
		return true
	} else {
		return false
	}
}

func (E ElasticSearchAction) IndexCreate(index string) bool {
	response, e := E.client.Indices.Create(strings.ToLower(index), func(request *esapi.IndicesCreateRequest) {
		request.Pretty = true
		request.Human = true
	})
	if e != nil {
		log.Println(e)
		return false
	}
	defer response.Body.Close()
	log.Println(response.String())
	return true
}

func (E ElasticSearchAction) IndexGetMapping(index string) (bool, string) {
	var response *esapi.Response
	var e error
	_, e = E.client.Indices.GetMapping(func(request *esapi.IndicesGetMappingRequest) {
		request = &esapi.IndicesGetMappingRequest{
			Index:  []string{index},
			Pretty: true,
			Human:  true,
		}
		response, e = request.Do(context.TODO(), E.client)
		if e != nil {
			log.Println(e)
			return
		}
	})
	if e != nil {
		log.Println(e)
		return false, "nil"
	}
	defer response.Body.Close()
	return true, response.String()

}

func (E ElasticSearchAction) IndexGetMappingField(index string, fields []string) (bool, string) {
	var response *esapi.Response
	var e error
	E.client.Indices.GetFieldMapping(fields, func(request *esapi.IndicesGetFieldMappingRequest) {
		request.Pretty = true
		request.Human = true
		request.Index = []string{index}
		response, e = request.Do(context.TODO(), E.client)
		if e != nil {
			log.Println(e)
			return
		}
	})
	if e != nil {
		return false, "nil"
	}
	return true, response.String()
}

func ExampleEsClient() {
	DefaultElasticSearchAction.IndexIsExists("users")
	ok := DefaultElasticSearchAction.IndexIsExists("python")
	fmt.Println(ok)
	if !ok {
		fmt.Println(DefaultElasticSearchAction.IndexCreate("Python"))
	}
	fmt.Println(DefaultElasticSearchAction.IndexGetMapping("python"))
	fmt.Println(DefaultElasticSearchAction.IndexGetMapping("golang"))
	fmt.Println(DefaultElasticSearchAction.IndexGetMappingField("users",[]string{"user","num"}))
}
