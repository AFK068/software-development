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

func TestAnalyticsFacade_CalculateBalanceDifference(t *testing.T) {
	mockRepo := mocks.NewRepository(t)
	facade := facades.NewAnalyticsFacade(mockRepo)

	operations := []*domain.Operation{
		{ID: 1, Type: domain.OperationTypeIncome, Amount: 1000},
		{ID: 2, Type: domain.OperationTypeExpense, Amount: 500},
		{ID: 3, Type: domain.OperationTypeIncome, Amount: 200},
	}

	mockRepo.On("GetAllOperations").Return(operations, nil)

	diff, err := facade.CalculateBalanceDifference()
	assert.NoError(t, err)
	assert.Equal(t, 700.0, diff)
	mockRepo.AssertExpectations(t)
}

func TestAnalyticsFacade_CalculateBalanceDifference_Error(t *testing.T) {
	mockRepo := mocks.NewRepository(t)
	facade := facades.NewAnalyticsFacade(mockRepo)

	mockRepo.On("GetAllOperations").Return(nil, errors.New("repository error"))

	_, err := facade.CalculateBalanceDifference()
	assert.Error(t, err)
	assert.EqualError(t, err, "repository error")
	mockRepo.AssertExpectations(t)
}

func TestAnalyticsFacade_CalculateBalance(t *testing.T) {
	mockRepo := mocks.NewRepository(t)
	facade := facades.NewAnalyticsFacade(mockRepo)

	accounts := []*domain.BankAccount{
		{ID: 1, Balance: 1000},
		{ID: 2, Balance: 2000},
	}

	mockRepo.On("GetAllBankAccounts").Return(accounts, nil)

	balance, err := facade.CalculateBalance()
	assert.NoError(t, err)
	assert.Equal(t, 3000.0, balance)
	mockRepo.AssertExpectations(t)
}

func TestAnalyticsFacade_CalculateBalance_Error(t *testing.T) {
	mockRepo := mocks.NewRepository(t)
	facade := facades.NewAnalyticsFacade(mockRepo)

	mockRepo.On("GetAllBankAccounts").Return(nil, errors.New("repository error"))

	_, err := facade.CalculateBalance()
	assert.Error(t, err)
	assert.EqualError(t, err, "repository error")
	mockRepo.AssertExpectations(t)
}

func TestAnalyticsFacade_RecalculateBalances(t *testing.T) {
	mockRepo := mocks.NewRepository(t)
	facade := facades.NewAnalyticsFacade(mockRepo)

	accounts := []*domain.BankAccount{
		{ID: 1, Balance: 0},
		{ID: 2, Balance: 0},
	}

	operationsAccount1 := []*domain.Operation{
		{ID: 1, Type: domain.OperationTypeIncome, Amount: 1000},
		{ID: 2, Type: domain.OperationTypeExpense, Amount: 500},
	}

	operationsAccount2 := []*domain.Operation{
		{ID: 3, Type: domain.OperationTypeIncome, Amount: 2000},
		{ID: 4, Type: domain.OperationTypeExpense, Amount: 1000},
	}

	mockRepo.On("GetAllBankAccounts").Return(accounts, nil)
	mockRepo.On("GetOperationsByAccountID", 1).Return(operationsAccount1, nil)
	mockRepo.On("GetOperationsByAccountID", 2).Return(operationsAccount2, nil)
	mockRepo.On("SaveBankAccount", mock.Anything).Return(nil)

	err := facade.RecalculateBalances()
	assert.NoError(t, err)

	assert.Equal(t, 500.0, accounts[0].Balance)
	assert.Equal(t, 1000.0, accounts[1].Balance)
	mockRepo.AssertExpectations(t)
}

func TestAnalyticsFacade_RecalculateBalances_Error(t *testing.T) {
	mockRepo := mocks.NewRepository(t)
	facade := facades.NewAnalyticsFacade(mockRepo)

	mockRepo.On("GetAllBankAccounts").Return(nil, errors.New("repository error"))

	err := facade.RecalculateBalances()
	assert.Error(t, err)
	assert.EqualError(t, err, "repository error")
	mockRepo.AssertExpectations(t)
}
