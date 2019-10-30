package main

import "encoding/json"

type SendMessage struct {
	Method  string `json:"method"`
	URL     string `json:"url"`
	Value   string `json:"value"`
	Date    string `json:"date"`
	encoded []byte
	err     error
}

func (S *SendMessage) Length() int {
	b, e := json.Marshal(S)
	S.encoded = b
	S.err = e
	return len(string(b))
}
func (S *SendMessage) Encode() ([]byte, error) {
	return S.encoded, S.err
}
