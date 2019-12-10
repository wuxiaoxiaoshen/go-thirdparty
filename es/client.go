package es

import (
	"github.com/elastic/go-elasticsearch/v7"
	"log"
)

type ClientESAction struct {
	Client *elasticsearch.Client
}

func NewEsClientAction() *ClientESAction {
	client,e := elasticsearch.NewDefaultClient()
	if e!=nil{
		log.Println("es client: fail")
		return nil
	}
	return &ClientESAction{Client:client}
}

func (C ClientESAction) Version()string{
	return elasticsearch.Version
}