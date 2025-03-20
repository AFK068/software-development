package parser

type DataParser interface {
	Parse(data []byte, v interface{}) error
}
