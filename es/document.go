package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/elastic/go-elasticsearch/v7/esutil"
	"log"
	"strings"
)

type Document struct {
	Index string `json:"index"`
}

var DefaultDocument = NewDocument("java")

func NewDocument(index string) *Document {
	return &Document{Index:strings.ToLower(index)}
}

type DocumentFields struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Hot string `json:"hot"`
	Excerpt string `json:"excerpt"`
}

func (D Document) Create(client *elasticsearch.Client, doc DocumentFields) (bool, string) {
	var (
		response *esapi.Response
		e error
	)
	client.Create(D.Index, doc.ID, esutil.NewJSONReader(doc), func(request *esapi.CreateRequest) {
		request.Pretty = true
		request.Human = true
		request.Body = esutil.NewJSONReader(doc)
		response,e = request.Do(context.TODO(), client)
	})
	if e!=nil{
		return false, "nil"
	}
	return true, response.String()

}
func (D Document) Get(index string,id string,client *elasticsearch.Client) (bool, string){
	var (
		response *esapi.Response
		e error
	)
	client.Get(index, id, func(request *esapi.GetRequest) {
		request.Human = true
		request.Pretty = true
		response,e = request.Do(context.TODO(), client)
	})
	if e!=nil{
		log.Println(e)
		return false, "nil"
	}
	defer response.Body.Close()
	return true, response.String()

}
func (D Document) FindAll() {}

func (D Document) InsertOne(ctx context.Context, client *elasticsearch.Client,doc DocumentFields)(bool,string) {
	m, _ := json.Marshal(doc)
	request := esapi.IndicesCreateRequest{
		Index: D.Index,
		Body: bytes.NewReader(m),
		Pretty: true,
		Human: true,
	}
	response, e := request.Do(ctx,client)
	if e!=nil{
		log.Println(e)
		return false,"nil"
	}
	defer response.Body.Close()
	return true, response.String()
}

func (D Document) BulkInsert(){}

func (D Document) SQLFind(){}

func (D Document) URLSearch(){}
func (D Document) ResponseBodySearch(){}
func (D Document) DeleteOne(){}

func ExampleDocument() {
	doc := DocumentFields{Title:"123",Hot:"34", Excerpt:"233", ID:"2"}
	if ok, content := DefaultDocument.Get("java", doc.ID, DefaultElasticSearchAction.client); !ok{
		fmt.Println(DefaultDocument.Create(DefaultElasticSearchAction.client, doc))
	}else{
		log.Println(content)
	}

}