package main

import "encoding/json"

type Message struct {
	Value  string `json:"value"`
	Number int    `json:"number"`
}

func (M *Message) Encoder() ([]byte, error) {
	return json.Marshal(M)
}
