package importer

import (
	"github.com/AFK068/bot/internal/domain"
	"github.com/AFK068/bot/internal/infrastructure/parser"
)

type OperationImporter struct {
	BaseImporter
	parser parser.DataParser
}

func NewOperationImporter(parser parser.DataParser) *OperationImporter {
	return &OperationImporter{
		parser: parser,
	}
}

func (o *OperationImporter) Import(fileName string) (interface{}, error) {
	data, err := o.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var result []*domain.Operation
	if err := o.parser.Parse(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}
