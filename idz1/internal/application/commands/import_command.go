package commands

import "github.com/AFK068/bot/internal/infrastructure/importer"

type Import struct {
	Importer importer.Importer
	FileName string
}

func NewImport(importer importer.Importer, fileName string) *Import {
	return &Import{
		Importer: importer,
		FileName: fileName,
	}
}

func (c *Import) Execute() error {
	_, err := c.Importer.Import(c.FileName)
	return err
}
