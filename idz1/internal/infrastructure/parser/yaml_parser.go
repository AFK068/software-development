package parser

import (
	"gopkg.in/yaml.v2"
)

type YAML struct {
}

func NewYAML() *YAML {
	return &YAML{}
}

func (y *YAML) Parse(data []byte, v interface{}) error {
	return yaml.Unmarshal(data, v)
}
