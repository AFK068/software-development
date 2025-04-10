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

func TestDefaultCategoryHandler_CreateCategory(t *testing.T) {
	mockFacade := mocks.NewCategoryService(t)
	mockUI := uiMocks.NewUserInput(t)

	handler := handler.NewCategoryHandler(mockFacade, mockUI)

	mockUI.On("GetIntValue", "Enter the category ID", 0, math.MaxInt).Return(1, nil)
	mockUI.On("GetStringValue", "Enter the category name").Return("Test Category", nil)
	mockUI.On("GetStringValue", "Enter the category type (e.g., income/expense)").Return("income", nil)

	mockFacade.On("CreateCategory", mock.MatchedBy(func(cat *domain.Category) bool {
		return cat.ID == 1 &&
			cat.Name == "Test Category" &&
			cat.Type == "income"
	})).Return(nil)

	err := handler.CreateCategory()
	assert.NoError(t, err)

	mockUI.AssertExpectations(t)
	mockFacade.AssertExpectations(t)
}

func TestDefaultCategoryHandler_UpdateCategory(t *testing.T) {
	mockFacade := mocks.NewCategoryService(t)
	mockUI := uiMocks.NewUserInput(t)

	handler := handler.NewCategoryHandler(mockFacade, mockUI)

	mockUI.On("GetIntValue", "Enter the category ID", 0, math.MaxInt).Return(1, nil)
	mockFacade.On("GetCategory", 1).Return(&domain.Category{
		ID:   1,
		Name: "Old Category",
		Type: "expense",
	}, nil)

	mockUI.On("GetStringValue", "Enter the category name").Return("Updated Category", nil)
	mockUI.On("GetStringValue", "Enter the category type (e.g., income/expense)").Return("income", nil)

	mockFacade.On("UpdateCategory", mock.MatchedBy(func(cat *domain.Category) bool {
		return cat.ID == 1 &&
			cat.Name == "Updated Category" &&
			cat.Type == "income"
	})).Return(nil)

	err := handler.UpdateCategory()
	assert.NoError(t, err)

	mockUI.AssertExpectations(t)
	mockFacade.AssertExpectations(t)
}

func TestDefaultCategoryHandler_DeleteCategory(t *testing.T) {
	mockFacade := mocks.NewCategoryService(t)
	mockUI := uiMocks.NewUserInput(t)

	handler := handler.NewCategoryHandler(mockFacade, mockUI)

	mockUI.On("GetIntValue", "Enter the category ID to delete", 0, math.MaxInt).Return(1, nil)

	mockFacade.On("DeleteCategory", 1).Return(nil)

	err := handler.DeleteCategory()
	assert.NoError(t, err)

	mockUI.AssertExpectations(t)
	mockFacade.AssertExpectations(t)
}
