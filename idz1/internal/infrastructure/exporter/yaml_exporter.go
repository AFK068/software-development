package exporter

import (
	"os"

	"gopkg.in/yaml.v2"

	"github.com/AFK068/bot/internal/domain"
)

type YAML struct {
	data []any
}

func NewYAML() *YAML {
	return &YAML{}
}

func (y *YAML) GetData() []any {
	return y.data
}

func (y *YAML) VisitBankAccount(account *domain.BankAccount) error {
	y.data = append(y.data, account)
	return nil
}

func (y *YAML) VisitCategory(category *domain.Category) error {
	y.data = append(y.data, category)
	return nil
}

func (y *YAML) VisitOperation(operation *domain.Operation) error {
	y.data = append(y.data, operation)
	return nil
}

func (y *YAML) Export(fileName string) error {
	data, err := yaml.Marshal(y.data)
	if err != nil {
		return err
	}

	err = os.WriteFile(fileName, data, 0o600)
	if err != nil {
		return err
	}

	return nil
}
