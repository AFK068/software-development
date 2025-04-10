package facades_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/AFK068/bot/internal/application/facades"
	"github.com/AFK068/bot/internal/domain"
	"github.com/AFK068/bot/internal/domain/mocks"
)

func TestOperationFacade_CreateOperation(t *testing.T) {
	mockRepo := mocks.NewRepository(t)
	facade := facades.NewOperationFacade(mockRepo)

	account := &domain.BankAccount{ID: 1, Balance: 1000}
	operation := &domain.Operation{ID: 1, Type: domain.OperationTypeIncome, Amount: 500, BankAccountID: 1}

	mockRepo.On("GetBankAccount", 1).Return(account, nil)
	mockRepo.On("SaveBankAccount", mock.Anything).Return(nil)
	mockRepo.On("SaveOperation", operation).Return(nil)

	err := facade.CreateOperation(operation)
	assert.NoError(t, err)

	assert.Equal(t, 1500.0, account.Balance)
	mockRepo.AssertExpectations(t)
}

func TestOperationFacade_CreateOperation_Expense(t *testing.T) {
	mockRepo := mocks.NewRepository(t)
	facade := facades.NewOperationFacade(mockRepo)

	account := &domain.BankAccount{ID: 1, Balance: 1000}
	operation := &domain.Operation{ID: 1, Type: domain.OperationTypeExpense, Amount: 200, BankAccountID: 1}

	mockRepo.On("GetBankAccount", 1).Return(account, nil)
	mockRepo.On("SaveBankAccount", mock.Anything).Return(nil)
	mockRepo.On("SaveOperation", operation).Return(nil)

	err := facade.CreateOperation(operation)
	assert.NoError(t, err)

	assert.Equal(t, 800.0, account.Balance)
	mockRepo.AssertExpectations(t)
}

func TestOperationFacade_CreateOperation_Error(t *testing.T) {
	mockRepo := mocks.NewRepository(t)
	facade := facades.NewOperationFacade(mockRepo)

	mockRepo.On("GetBankAccount", 1).Return(nil, errors.New("account not found"))

	operation := &domain.Operation{ID: 1, Type: domain.OperationTypeIncome, Amount: 500, BankAccountID: 1}

	err := facade.CreateOperation(operation)
	assert.Error(t, err)
	assert.EqualError(t, err, "account not found")
	mockRepo.AssertExpectations(t)
}

func TestOperationFacade_GetOperation(t *testing.T) {
	mockRepo := mocks.NewRepository(t)
	facade := facades.NewOperationFacade(mockRepo)

	operation := &domain.Operation{ID: 1, Type: domain.OperationTypeIncome, Amount: 500}

	mockRepo.On("GetOperation", 1).Return(operation, nil)

	result, err := facade.GetOperation(1)
	assert.NoError(t, err)
	assert.Equal(t, operation, result)
	mockRepo.AssertExpectations(t)
}

func TestOperationFacade_DeleteOperation(t *testing.T) {
	mockRepo := mocks.NewRepository(t)
	facade := facades.NewOperationFacade(mockRepo)

	mockRepo.On("DeleteOperation", 1).Return(nil)

	err := facade.DeleteOperation(1)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestOperationFacade_UpdateOperation(t *testing.T) {
	mockRepo := mocks.NewRepository(t)
	facade := facades.NewOperationFacade(mockRepo)

	operation := &domain.Operation{ID: 1, Type: domain.OperationTypeIncome, Amount: 500}

	mockRepo.On("UpdateOperation", operation).Return(nil)

	err := facade.UpdateOperation(operation)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
