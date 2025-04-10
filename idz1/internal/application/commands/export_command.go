package commands

import "github.com/AFK068/bot/internal/infrastructure/exporter"

type Export struct {
	Exporter exporter.Exporter
	FileName string
}

func NewExport(exporter exporter.Exporter, fileName string) *Export {
	return &Export{
		Exporter: exporter,
		FileName: fileName,
	}
}

func (c *Export) Execute() error {
	return c.Exporter.Export(c.FileName)
}
