package exporter_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/AFK068/bot/internal/domain"
	"github.com/AFK068/bot/internal/infrastructure/exporter"
)

func TestYAML_VisitBankAccount(t *testing.T) {
	yamlExporter := exporter.NewYAML()

	bankAccount := &domain.BankAccount{ID: 1, Name: "Test Account", Balance: 1000}

	err := yamlExporter.VisitBankAccount(bankAccount)
	assert.NoError(t, err)
	assert.Len(t, yamlExporter.GetData(), 1)
	assert.Equal(t, bankAccount, yamlExporter.GetData()[0])
}

func TestYAML_VisitCategory(t *testing.T) {
	yamlExporter := exporter.NewYAML()

	category := &domain.Category{ID: 1, Name: "Test Category", Type: "Expense"}

	err := yamlExporter.VisitCategory(category)
	assert.NoError(t, err)
	assert.Len(t, yamlExporter.GetData(), 1)
	assert.Equal(t, category, yamlExporter.GetData()[0])
}

func TestYAML_VisitOperation(t *testing.T) {
	yamlExporter := exporter.NewYAML()

	operation := &domain.Operation{ID: 1, Type: "Income", BankAccountID: 1, Amount: 500, Description: "Test Operation"}

	err := yamlExporter.VisitOperation(operation)
	assert.NoError(t, err)
	assert.Len(t, yamlExporter.GetData(), 1)
	assert.Equal(t, operation, yamlExporter.GetData()[0])
}

func TestYAML_Export(t *testing.T) {
	yamlExporter := exporter.NewYAML()

	bankAccount := &domain.BankAccount{ID: 1, Name: "Test Account", Balance: 1000}
	category := &domain.Category{ID: 1, Name: "Test Category", Type: "Expense"}
	operation := &domain.Operation{ID: 1, Type: "Income", BankAccountID: 1, Amount: 500, Description: "Test Operation"}

	assert.NoError(t, yamlExporter.VisitBankAccount(bankAccount))
	assert.NoError(t, yamlExporter.VisitCategory(category))
	assert.NoError(t, yamlExporter.VisitOperation(operation))

	fileName := "test_output.yaml"
	defer os.Remove(fileName)

	err := yamlExporter.Export(fileName)
	assert.NoError(t, err)

	data, err := os.ReadFile(fileName)
	assert.NoError(t, err)
	assert.Contains(t, string(data), "id: 1")
	assert.Contains(t, string(data), "name: Test Account")
	assert.Contains(t, string(data), "balance: 1000")
	assert.Contains(t, string(data), "type: Expense")
	assert.Contains(t, string(data), "amount: 500")
	assert.Contains(t, string(data), "description: Test Operation")
}
