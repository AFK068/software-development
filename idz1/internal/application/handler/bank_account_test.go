package handler_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/AFK068/bot/internal/application/facades/mocks"
	"github.com/AFK068/bot/internal/application/handler"
	"github.com/AFK068/bot/internal/domain"

	uiMocks "github.com/AFK068/bot/internal/infrastructure/userinteraction/mocks"
)

func TestDefaultBankAccountHandler_CreateBankAccount(t *testing.T) {
	mockFacade := mocks.NewBankAccountService(t)
	mockUI := uiMocks.NewUserInput(t)

	handler := handler.NewBankAccountHandler(mockFacade, mockUI)

	mockUI.On("GetIntValue", "Enter the account ID", 0, math.MaxInt).Return(1, nil)
	mockUI.On("GetStringValue", "Enter the account name").Return("Test Account", nil)
	mockUI.On("GetFloatValue", "Enter the account balance", 0.0, math.MaxFloat64).Return(1000.0, nil)

	mockFacade.On("CreateBankAccount", mock.MatchedBy(func(account *domain.BankAccount) bool {
		return account.ID == 1 &&
			account.Name == "Test Account" &&
			account.Balance == 1000.0
	})).Return(nil)

	err := handler.CreateBankAccount()
	assert.NoError(t, err)

	mockUI.AssertExpectations(t)
	mockFacade.AssertExpectations(t)
}

func TestDefaultBankAccountHandler_UpdateBankAccount(t *testing.T) {
	mockFacade := mocks.NewBankAccountService(t)
	mockUI := uiMocks.NewUserInput(t)

	handler := handler.NewBankAccountHandler(mockFacade, mockUI)

	mockUI.On("GetIntValue", "Enter the account ID", 0, math.MaxInt).Return(1, nil)
	mockFacade.On("GetBankAccount", 1).Return(&domain.BankAccount{
		ID:      1,
		Name:    "Old Account",
		Balance: 500.0,
	}, nil)

	mockUI.On("GetStringValue", "Enter the account name").Return("Updated Account", nil)
	mockUI.On("GetFloatValue", "Enter the account balance", 0.0, math.MaxFloat64).Return(2000.0, nil)

	mockFacade.On("UpdateBankAccount", mock.MatchedBy(func(account *domain.BankAccount) bool {
		return account.ID == 1 &&
			account.Name == "Updated Account" &&
			account.Balance == 2000.0
	})).Return(nil)

	err := handler.UpdateBankAccount()
	assert.NoError(t, err)

	mockUI.AssertExpectations(t)
	mockFacade.AssertExpectations(t)
}

func TestDefaultBankAccountHandler_DeleteBankAccount(t *testing.T) {
	mockFacade := mocks.NewBankAccountService(t)
	mockUI := uiMocks.NewUserInput(t)

	handler := handler.NewBankAccountHandler(mockFacade, mockUI)

	mockUI.On("GetIntValue", "Enter the account ID to delete", 0, math.MaxInt).Return(1, nil)

	mockFacade.On("DeleteBankAccount", 1).Return(nil)

	err := handler.DeleteBankAccount()
	assert.NoError(t, err)

	mockUI.AssertExpectations(t)
	mockFacade.AssertExpectations(t)
}
