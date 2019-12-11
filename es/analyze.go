package main

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"log"
	"strings"
)

type AnalyzeInfo struct {
	client *elasticsearch.Client
}

func (A AnalyzeInfo) Standard(index string, text string) {
	body := `
		{
			"analyzer": "standard",
			"text": ["`+text+`"]
		}
`
	response, e := A.client.Indices.Analyze(
		A.client.Indices.Analyze.WithIndex(index),
		A.client.Indices.Analyze.WithBody(strings.NewReader(body)),
		A.client.Indices.Analyze.WithPretty(),
		)
	if e!=nil{
		log.Println(e)
		return
	}
	fmt.Println(response.String())
}

func ExampleAnalyze() {
	analyzer := AnalyzeInfo{client:DefaultElasticSearchAction.client}
	analyzer.Standard("java", "Hello World")

}