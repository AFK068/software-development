package importer

type Importer interface {
	Import(fileName string) (interface{}, error)
}
