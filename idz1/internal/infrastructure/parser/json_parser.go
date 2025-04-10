package parser

import (
	"encoding/json"
)

type JSON struct {
}

func NewJSON() *JSON {
	return &JSON{}
}

func (j *JSON) Parse(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
