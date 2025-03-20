package importer

import (
	"io"
	"os"
)

type BaseImporter struct{}

func NewBaseImporter() *BaseImporter {
	return &BaseImporter{}
}

func (b *BaseImporter) ReadFile(fileName string) ([]byte, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return io.ReadAll(file)
}
