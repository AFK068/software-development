package importer

import (
	"github.com/AFK068/bot/internal/domain"
	"github.com/AFK068/bot/internal/infrastructure/parser"
)

type BankAccountImporter struct {
	BaseImporter
	parser parser.DataParser
}

func NewBankAccountImporter(parser parser.DataParser) *BankAccountImporter {
	return &BankAccountImporter{
		parser: parser,
	}
}

func (b *BankAccountImporter) Import(fileName string) (interface{}, error) {
	data, err := b.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var bankAccounts []*domain.BankAccount

	err = b.parser.Parse(data, &bankAccounts)
	if err != nil {
		return nil, err
	}

	return bankAccounts, nil
}
