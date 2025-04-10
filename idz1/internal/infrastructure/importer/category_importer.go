package importer

import (
	"github.com/AFK068/bot/internal/domain"
	"github.com/AFK068/bot/internal/infrastructure/parser"
)

type CategoryImporter struct {
	BaseImporter
	parser parser.DataParser
}

func NewCategoryImporter(parser parser.DataParser) *CategoryImporter {
	return &CategoryImporter{
		parser: parser,
	}
}

func (c *CategoryImporter) Import(fileName string) (interface{}, error) {
	data, err := c.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var result []*domain.Category
	if err := c.parser.Parse(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}
