package exporter_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/AFK068/bot/internal/domain"
	"github.com/AFK068/bot/internal/infrastructure/exporter"
)

func TestJSON_VisitBankAccount(t *testing.T) {
	jsonExporter := exporter.NewJSON()

	bankAccount := &domain.BankAccount{ID: 1, Name: "Test Account", Balance: 1000}

	err := jsonExporter.VisitBankAccount(bankAccount)
	assert.NoError(t, err)
	assert.Len(t, jsonExporter.GetData(), 1)
	assert.Equal(t, bankAccount, jsonExporter.GetData()[0])
}

func TestJSON_VisitCategory(t *testing.T) {
	jsonExporter := exporter.NewJSON()

	category := &domain.Category{ID: 1, Name: "Test Category", Type: "Expense"}

	err := jsonExporter.VisitCategory(category)
	assert.NoError(t, err)
	assert.Len(t, jsonExporter.GetData(), 1)
	assert.Equal(t, category, jsonExporter.GetData()[0])
}

func TestJSON_VisitOperation(t *testing.T) {
	jsonExporter := exporter.NewJSON()

	operation := &domain.Operation{ID: 1, Type: "Income", BankAccountID: 1, Amount: 500, Description: "Test Operation"}

	err := jsonExporter.VisitOperation(operation)
	assert.NoError(t, err)
	assert.Len(t, jsonExporter.GetData(), 1)
	assert.Equal(t, operation, jsonExporter.GetData()[0])
}

func TestJSON_Export(t *testing.T) {
	jsonExporter := exporter.NewJSON()

	bankAccount := &domain.BankAccount{ID: 1, Name: "Test Account", Balance: 1000}
	category := &domain.Category{ID: 1, Name: "Test Category", Type: "Expense"}
	operation := &domain.Operation{ID: 1, Type: "Income", BankAccountID: 1, Amount: 500, Description: "Test Operation"}

	assert.NoError(t, jsonExporter.VisitBankAccount(bankAccount))
	assert.NoError(t, jsonExporter.VisitCategory(category))
	assert.NoError(t, jsonExporter.VisitOperation(operation))

	fileName := "test_output.json"
	defer os.Remove(fileName)

	err := jsonExporter.Export(fileName)
	assert.NoError(t, err)

	data, err := os.ReadFile(fileName)
	assert.NoError(t, err)
	assert.Contains(t, string(data), `"id": 1`)
	assert.Contains(t, string(data), `"name": "Test Account"`)
	assert.Contains(t, string(data), `"balance": 1000`)
	assert.Contains(t, string(data), `"type": "Expense"`)
	assert.Contains(t, string(data), `"amount": 500`)
	assert.Contains(t, string(data), `"description": "Test Operation"`)
}
