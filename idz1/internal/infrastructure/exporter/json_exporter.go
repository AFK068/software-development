package exporter

import (
	"encoding/json"
	"os"

	"github.com/AFK068/bot/internal/domain"
)

type JSON struct {
	data []any
}

func NewJSON() *JSON {
	return &JSON{}
}

func (j *JSON) GetData() []any {
	return j.data
}

func (j *JSON) VisitBankAccount(account *domain.BankAccount) error {
	j.data = append(j.data, account)
	return nil
}

func (j *JSON) VisitCategory(category *domain.Category) error {
	j.data = append(j.data, category)
	return nil
}

func (j *JSON) VisitOperation(operation *domain.Operation) error {
	j.data = append(j.data, operation)
	return nil
}

func (j *JSON) Export(fileName string) error {
	data, err := json.MarshalIndent(j.data, "", " ")
	if err != nil {
		return nil
	}

	err = os.WriteFile(fileName, data, 0o600)
	if err != nil {
		return err
	}

	return nil
}
