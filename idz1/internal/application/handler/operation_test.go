package handler_test

import (
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/AFK068/bot/internal/application/facades/mocks"
	"github.com/AFK068/bot/internal/application/handler"
	"github.com/AFK068/bot/internal/domain"

	uiMocks "github.com/AFK068/bot/internal/infrastructure/userinteraction/mocks"
)

func TestDefaultOperationHandler_CreateOperation(t *testing.T) {
	mockFacade := mocks.NewOperationService(t)
	mockUI := uiMocks.NewUserInput(t)

	handler := handler.NewOperationHandler(mockFacade, mockUI)

	mockUI.On("GetIntValue", "Enter the ID of the operation", 0, math.MaxInt).Return(1, nil)
	mockUI.On("GetStringValue", "Enter the type of the operation (income/expense)").Return(domain.OperationTypeIncome, nil)
	mockUI.On("GetIntValue", "Enter the bank account ID", 0, math.MaxInt).Return(1, nil)
	mockUI.On("GetFloatValue", "Enter the amount of the operation", float64(0), math.MaxFloat64).Return(1000.0, nil)
	mockUI.On("GetTimeValue", "Enter the date of the operation (YYYY-MM-DD HH:MM:SS)").Return(time.Now(), nil)
	mockUI.On("GetStringValue", "Enter a description for the operation").Return("Test Operation", nil)
	mockUI.On("GetIntValue", "Enter the category ID", 0, math.MaxInt).Return(1, nil)

	mockFacade.On("CreateOperation", mock.MatchedBy(func(op *domain.Operation) bool {
		return op.ID == 1 &&
			op.Type == domain.OperationTypeIncome &&
			op.BankAccountID == 1 &&
			op.Amount == 1000.0 &&
			op.Description == "Test Operation" &&
			op.CategoryID == 1
	})).Return(nil)

	err := handler.CreateOperation()
	assert.NoError(t, err)

	mockUI.AssertExpectations(t)
	mockFacade.AssertExpectations(t)
}

func TestDefaultOperationHandler_UpdateOperation(t *testing.T) {
	mockFacade := mocks.NewOperationService(t)
	mockUI := uiMocks.NewUserInput(t)

	handler := handler.NewOperationHandler(mockFacade, mockUI)

	mockUI.On("GetIntValue", "Enter the ID of the operation", 0, math.MaxInt).Return(1, nil)
	mockFacade.On("GetOperation", 1).Return(&domain.Operation{
		ID:            1,
		Type:          domain.OperationTypeIncome,
		BankAccountID: 1,
		Amount:        500.0,
		Description:   "Old Operation",
		CategoryID:    1,
	}, nil)

	mockUI.On("GetStringValue", "Enter the type of the operation (income/expense)").Return(domain.OperationTypeExpense, nil)
	mockUI.On("GetIntValue", "Enter the bank account ID", 0, math.MaxInt).Return(1, nil)
	mockUI.On("GetFloatValue", "Enter the amount of the operation", float64(0), math.MaxFloat64).Return(1000.0, nil)
	mockUI.On("GetTimeValue", "Enter the date of the operation (YYYY-MM-DD HH:MM:SS)").Return(time.Now(), nil)
	mockUI.On("GetStringValue", "Enter a description for the operation").Return("Updated Operation", nil)
	mockUI.On("GetIntValue", "Enter the category ID", 0, math.MaxInt).Return(2, nil)

	mockFacade.On("UpdateOperation", mock.MatchedBy(func(op *domain.Operation) bool {
		return op.ID == 1 &&
			op.Type == domain.OperationTypeExpense &&
			op.BankAccountID == 1 &&
			op.Amount == 1000.0 &&
			op.Description == "Updated Operation" &&
			op.CategoryID == 2
	})).Return(nil)

	err := handler.UpdateOperation()
	assert.NoError(t, err)

	mockUI.AssertExpectations(t)
	mockFacade.AssertExpectations(t)
}

func TestDefaultOperationHandler_DeleteOperation(t *testing.T) {
	mockFacade := mocks.NewOperationService(t)
	mockUI := uiMocks.NewUserInput(t)

	handler := handler.NewOperationHandler(mockFacade, mockUI)

	mockUI.On("GetIntValue", "Enter the ID of the operation", 0, math.MaxInt).Return(1, nil)

	mockFacade.On("DeleteOperation", 1).Return(nil)

	err := handler.DeleteOperation()
	assert.NoError(t, err)

	mockUI.AssertExpectations(t)
	mockFacade.AssertExpectations(t)
}
