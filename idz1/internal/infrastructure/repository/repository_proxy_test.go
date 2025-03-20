package repository_test

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/AFK068/bot/internal/domain"
	"github.com/AFK068/bot/internal/domain/mocks"
	"github.com/AFK068/bot/internal/infrastructure/repository"
)

func TestProxyRepo_SaveBankAccount(t *testing.T) {
	mockRepo := mocks.NewRepository(t)
	proxyRepo := repository.NewProxy(mockRepo, 5*time.Minute, 10*time.Minute)

	account := &domain.BankAccount{ID: 1, Name: "Test Account", Balance: 1000}

	mockRepo.On("SaveBankAccount", account).Return(nil)

	err := proxyRepo.SaveBankAccount(account)
	assert.NoError(t, err)
}

func TestProxyRepo_GetBankAccount(t *testing.T) {
	mockRepo := mocks.NewRepository(t)
	proxyRepo := repository.NewProxy(mockRepo, 5*time.Minute, 10*time.Minute)

	account := &domain.BankAccount{ID: 1, Name: "Test Account", Balance: 1000}

	mockRepo.On("GetBankAccount", 1).Return(account, nil)

	retrievedAccount, err := proxyRepo.GetBankAccount(1)
	assert.NoError(t, err)
	assert.Equal(t, account, retrievedAccount)

	retrievedAccountFromCache, err := proxyRepo.GetBankAccount(1)
	assert.NoError(t, err)
	assert.Equal(t, account, retrievedAccountFromCache)
}

func TestProxyRepo_DeleteBankAccount(t *testing.T) {
	mockRepo := mocks.NewRepository(t)
	proxyRepo := repository.NewProxy(mockRepo, 5*time.Minute, 10*time.Minute)

	account := &domain.BankAccount{ID: 1, Name: "Test Account", Balance: 1000}
	mockRepo.On("SaveBankAccount", account).Return(nil)
	err := proxyRepo.SaveBankAccount(account)
	assert.NoError(t, err)

	mockRepo.On("DeleteBankAccount", 1).Return(nil)
	mockRepo.On("GetBankAccount", 1).Return(nil, errors.New("bank account not found"))

	err = proxyRepo.DeleteBankAccount(1)
	assert.NoError(t, err)

	_, err = proxyRepo.GetBankAccount(1)
	assert.Error(t, err)
	assert.EqualError(t, err, "bank account not found")
}

func TestProxyRepo_UpdateBankAccount(t *testing.T) {
	mockRepo := mocks.NewRepository(t)
	proxyRepo := repository.NewProxy(mockRepo, 5*time.Minute, 10*time.Minute)

	account := &domain.BankAccount{ID: 1, Name: "Test Account", Balance: 1000}

	mockRepo.On("UpdateBankAccount", account).Return(nil)

	err := proxyRepo.UpdateBankAccount(account)
	assert.NoError(t, err)
}

func TestProxyRepo_GetAllBankAccounts(t *testing.T) {
	mockRepo := mocks.NewRepository(t)
	proxyRepo := repository.NewProxy(mockRepo, 5*time.Minute, 10*time.Minute)

	accounts := []*domain.BankAccount{
		{ID: 1, Name: "Account 1", Balance: 1000},
		{ID: 2, Name: "Account 2", Balance: 2000},
	}

	mockRepo.On("GetAllBankAccounts").Return(accounts, nil)

	retrievedAccounts, err := proxyRepo.GetAllBankAccounts()
	assert.NoError(t, err)
	assert.Equal(t, accounts, retrievedAccounts)
}

func TestProxyRepo_SaveCategory(t *testing.T) {
	mockRepo := mocks.NewRepository(t)
	proxyRepo := repository.NewProxy(mockRepo, 5*time.Minute, 10*time.Minute)

	category := &domain.Category{ID: 1, Name: "Test Category", Type: "Expense"}

	mockRepo.On("SaveCategory", category).Return(nil)

	err := proxyRepo.SaveCategory(category)
	assert.NoError(t, err)
}

func TestProxyRepo_GetCategory(t *testing.T) {
	mockRepo := mocks.NewRepository(t)
	proxyRepo := repository.NewProxy(mockRepo, 5*time.Minute, 10*time.Minute)

	category := &domain.Category{ID: 1, Name: "Test Category", Type: "Expense"}

	mockRepo.On("GetCategory", 1).Return(category, nil)

	retrievedCategory, err := proxyRepo.GetCategory(1)
	assert.NoError(t, err)
	assert.Equal(t, category, retrievedCategory)

	retrievedCategoryFromCache, err := proxyRepo.GetCategory(1)
	assert.NoError(t, err)
	assert.Equal(t, category, retrievedCategoryFromCache)
}

func TestProxyRepo_DeleteCategory(t *testing.T) {
	mockRepo := mocks.NewRepository(t)
	proxyRepo := repository.NewProxy(mockRepo, 5*time.Minute, 10*time.Minute)

	category := &domain.Category{ID: 1, Name: "Test Category", Type: "Expense"}
	mockRepo.On("SaveCategory", category).Return(nil)
	err := proxyRepo.SaveCategory(category)
	assert.NoError(t, err)

	mockRepo.On("DeleteCategory", 1).Return(nil)
	mockRepo.On("GetCategory", 1).Return(nil, errors.New("category not found"))

	err = proxyRepo.DeleteCategory(1)
	assert.NoError(t, err)

	_, err = proxyRepo.GetCategory(1)
	assert.Error(t, err)
	assert.EqualError(t, err, "category not found")
}

func TestProxyRepo_UpdateCategory(t *testing.T) {
	mockRepo := mocks.NewRepository(t)
	proxyRepo := repository.NewProxy(mockRepo, 5*time.Minute, 10*time.Minute)

	category := &domain.Category{ID: 1, Name: "Test Category", Type: "Expense"}

	mockRepo.On("UpdateCategory", category).Return(nil)

	err := proxyRepo.UpdateCategory(category)
	assert.NoError(t, err)
}

func TestProxyRepo_GetAllCategories(t *testing.T) {
	mockRepo := mocks.NewRepository(t)
	proxyRepo := repository.NewProxy(mockRepo, 5*time.Minute, 10*time.Minute)

	categories := []*domain.Category{
		{ID: 1, Name: "Category 1", Type: "Expense"},
		{ID: 2, Name: "Category 2", Type: "Income"},
	}

	mockRepo.On("GetAllCategories").Return(categories, nil)

	retrievedCategories, err := proxyRepo.GetAllCategories()
	assert.NoError(t, err)
	assert.Equal(t, categories, retrievedCategories)
}

func TestProxyRepo_SaveOperation(t *testing.T) {
	mockRepo := mocks.NewRepository(t)
	proxyRepo := repository.NewProxy(mockRepo, 5*time.Minute, 10*time.Minute)

	operation := &domain.Operation{ID: 1, Type: "Income", Amount: 1000}

	mockRepo.On("SaveOperation", operation).Return(nil)

	err := proxyRepo.SaveOperation(operation)
	assert.NoError(t, err)
}

func TestProxyRepo_GetOperation(t *testing.T) {
	mockRepo := mocks.NewRepository(t)
	proxyRepo := repository.NewProxy(mockRepo, 5*time.Minute, 10*time.Minute)

	operation := &domain.Operation{ID: 1, Type: "Income", Amount: 1000}

	mockRepo.On("GetOperation", 1).Return(operation, nil)

	retrievedOperation, err := proxyRepo.GetOperation(1)
	assert.NoError(t, err)
	assert.Equal(t, operation, retrievedOperation)

	retrievedOperationFromCache, err := proxyRepo.GetOperation(1)
	assert.NoError(t, err)
	assert.Equal(t, operation, retrievedOperationFromCache)
}

func TestProxyRepo_DeleteOperation(t *testing.T) {
	mockRepo := mocks.NewRepository(t)
	proxyRepo := repository.NewProxy(mockRepo, 5*time.Minute, 10*time.Minute)

	operation := &domain.Operation{ID: 1, Type: "Income", Amount: 1000}
	mockRepo.On("SaveOperation", operation).Return(nil)
	err := proxyRepo.SaveOperation(operation)
	assert.NoError(t, err)

	mockRepo.On("DeleteOperation", 1).Return(nil)
	mockRepo.On("GetOperation", 1).Return(nil, errors.New("operation not found"))

	err = proxyRepo.DeleteOperation(1)
	assert.NoError(t, err)

	_, err = proxyRepo.GetOperation(1)
	assert.Error(t, err)
	assert.EqualError(t, err, "operation not found")
}

func TestProxyRepo_UpdateOperation(t *testing.T) {
	mockRepo := mocks.NewRepository(t)
	proxyRepo := repository.NewProxy(mockRepo, 5*time.Minute, 10*time.Minute)

	operation := &domain.Operation{ID: 1, Type: "Income", Amount: 1000}

	mockRepo.On("UpdateOperation", operation).Return(nil)

	err := proxyRepo.UpdateOperation(operation)
	assert.NoError(t, err)
}

func TestProxyRepo_GetAllOperations(t *testing.T) {
	mockRepo := mocks.NewRepository(t)
	proxyRepo := repository.NewProxy(mockRepo, 5*time.Minute, 10*time.Minute)

	operations := []*domain.Operation{
		{ID: 1, Type: "Income", Amount: 1000},
		{ID: 2, Type: "Expense", Amount: 500},
	}

	mockRepo.On("GetAllOperations").Return(operations, nil)

	retrievedOperations, err := proxyRepo.GetAllOperations()
	assert.NoError(t, err)
	assert.Equal(t, operations, retrievedOperations)
}
