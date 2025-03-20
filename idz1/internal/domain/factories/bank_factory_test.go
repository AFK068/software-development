package factories_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/AFK068/bot/internal/domain"
	"github.com/AFK068/bot/internal/domain/factories"
)

func TestBankFactory_NewBankAccount_Success(t *testing.T) {
	factory := factories.NewBankFactory()

	account, err := factory.NewBankAccount(1, "Test Account", 1000)
	assert.NoError(t, err)
	assert.NotNil(t, account)
	assert.Equal(t, 1, account.ID)
	assert.Equal(t, "Test Account", account.Name)
	assert.Equal(t, 1000.0, account.Balance)
}

func TestBankFactory_NewBankAccount_Failure(t *testing.T) {
	factory := factories.NewBankFactory()

	_, err := factory.NewBankAccount(2, "", 1000)
	assert.Error(t, err)
	assert.EqualError(t, err, "name is required")

	_, err = factory.NewBankAccount(3, "Invalid Account", -500)
	assert.Error(t, err)
	assert.EqualError(t, err, "negative balance")
}

func TestBankFactory_NewCategory_Success(t *testing.T) {
	factory := factories.NewBankFactory()

	category, err := factory.NewCategory(1, "Test Category", "Expense")
	assert.NoError(t, err)
	assert.NotNil(t, category)
	assert.Equal(t, 1, category.ID)
	assert.Equal(t, "Test Category", category.Name)
	assert.Equal(t, "Expense", category.Type)
}

func TestBankFactory_NewCategory_Failure(t *testing.T) {
	factory := factories.NewBankFactory()

	_, err := factory.NewCategory(2, "", "Expense")
	assert.Error(t, err)
	assert.EqualError(t, err, "name is required")

	_, err = factory.NewCategory(3, "Invalid Category", "")
	assert.Error(t, err)
	assert.EqualError(t, err, "category type is required")
}

func TestBankFactory_NewOperation_Success(t *testing.T) {
	factory := factories.NewBankFactory()

	date := time.Now()
	operation, err := factory.NewOperation(1, domain.OperationTypeIncome, 1, 1000, date, "Test Operation", 1)
	assert.NoError(t, err)
	assert.NotNil(t, operation)
	assert.Equal(t, 1, operation.ID)
	assert.Equal(t, domain.OperationTypeIncome, operation.Type)
	assert.Equal(t, 1, operation.BankAccountID)
	assert.Equal(t, 1000.0, operation.Amount)
	assert.Equal(t, date, operation.Date)
	assert.Equal(t, "Test Operation", operation.Description)
	assert.Equal(t, 1, operation.CategoryID)
}

func TestBankFactory_NewOperation_Failure(t *testing.T) {
	factory := factories.NewBankFactory()

	date := time.Now()

	_, err := factory.NewOperation(2, "", 1, 1000, date, "Test Operation", 1)
	assert.Error(t, err)
	assert.EqualError(t, err, "type is required")

	_, err = factory.NewOperation(3, domain.OperationTypeExpense, 1, -1000, date, "Test Operation", 1)
	assert.Error(t, err)
	assert.EqualError(t, err, "negative amount")

	_, err = factory.NewOperation(4, domain.OperationTypeIncome, 1, 1000, time.Time{}, "Test Operation", 1)
	assert.Error(t, err)
	assert.EqualError(t, err, "date is required")

	_, err = factory.NewOperation(5, domain.OperationTypeIncome, 1, 1000, date, "", 1)
	assert.Error(t, err)
	assert.EqualError(t, err, "description is required")
}
