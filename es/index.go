package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"log"
	"strings"
)

type Index struct {
	name string
}

type MappingSettings struct {
	Mapping AllField `json:"mapping"`
}

type AllField struct {
	Properties []Field `json:"properties"`
}
type Field struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func NewIndex(name string) *Index{
	return &Index{name:strings.ToLower(name)}
}

func (I Index) Exists(ctx context.Context, client *elasticsearch.Client) bool {
	request := esapi.IndicesExistsRequest{
		Index: []string{I.name},
		Pretty:true,
	    Human:true,
	}
	response, e := request.Do(ctx, client)
	if e!=nil{
		log.Println("index exists", e)
		return false
	}
	defer response.Body.Close()
	if response.StatusCode!=200 {
		return false
	}
	log.Println("?",response.String())
	return true
}

func (I Index) Create(ctx context.Context, client *elasticsearch.Client) (bool, string) {
	// mapping 的作用相当于 schema, 主要是对索引字段的定义：包括类型，是否建立索引
	mapping := map[string]interface{}{
		"mappings": map[string]interface{}{
				"properties": map[string]interface{}{
					"str": map[string]interface{}{
						"type": "keyword",
					},
					"name": map[string]interface{}{
					    "type": "text",
					},
				},
			},
	}
	m, _ := json.Marshal(mapping)
	request := esapi.IndicesCreateRequest{
		Index: I.name,
		Body: bytes.NewReader(m),
	}
	response, e := request.Do(ctx, client)
	if e!=nil{
		log.Println(e)
		return false, "nil"
	}
	defer response.Body.Close()
	return true, response.String()

}

func (I Index) UpdateMapping(ctx context.Context, client *elasticsearch.Client, field Field) (bool, string) {
	upgrade := map[string]interface{}{
		"properties": map[string] interface{}{
			field.Name: map[string]interface{}{
				"type": field.Type,
			},
		},
	}
	m, _ := json.Marshal(upgrade)
	request := esapi.IndicesPutMappingRequest{
		Index: []string{I.name},
		Body: bytes.NewReader(m),
		Pretty: true,
		Human: true,
	}
	response, e := request.Do(ctx, client)
	if e!=nil{
		log.Println(e)
		return false, "nil"
	}
	defer response.Body.Close()
	return true, response.String()
}

func (I Index) GetMapping(ctx context.Context, client *elasticsearch.Client) (bool, string) {
	request := esapi.IndicesGetMappingRequest{
		Index:[]string{I.name},
		Pretty: true,
		Human: true,
	}
	response, e := request.Do(ctx, client)
	if e!=nil{
		log.Println(e)
		return false, "nil"
	}
	defer response.Body.Close()
	return true, response.String()
}

func (I Index) GetMappingField(ctx context.Context, client *elasticsearch.Client, fields []string)(bool, string) {

	request := esapi.IndicesGetFieldMappingRequest{
		Index: []string{I.name},
		Fields:fields,
		Pretty: true,
		Human: true,
	}
	response, e := request.Do(ctx, client)
	if e!=nil{
		log.Println(e)
		return false, "nil"
	}
	defer response.Body.Close()
	return true, response.String()

}

func (I Index) DeleteIndex(ctx context.Context, client *elasticsearch.Client) bool {
	request := esapi.IndicesDeleteRequest{
		Index:[]string{I.name},
		Pretty:true,
		Human:true,
	}
	response,e := request.Do(ctx, client)
	if e!=nil{
		log.Println(e)
		return false
	}
	defer response.Body.Close()
	return true
}

func ExampleForIndex() {
	es1 := DefaultElasticSearchAction
	index:= NewIndex("Golang")
	ok := index.Exists(context.TODO(), es1.client)
	if !ok {
		yes, content := index.Create(context.TODO(), es1.client)
		fmt.Println(content)
		if yes {
			fmt.Println(index.GetMapping(context.TODO(), es1.client))
		}

	}
	filed := Field{Type:"integer",Name: "age"}
	yes, content := index.UpdateMapping(context.TODO(), es1.client, filed)
	fmt.Println(yes, content)
	fmt.Println(index.GetMappingField(context.TODO(), es1.client, []string{"age", "name"}))

}