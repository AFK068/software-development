package facades_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/AFK068/bot/internal/application/facades"
	"github.com/AFK068/bot/internal/domain"
	"github.com/AFK068/bot/internal/domain/mocks"
)

func TestCategoryFacade_CreateCategory(t *testing.T) {
	mockRepo := mocks.NewRepository(t)
	facade := facades.NewCategoryFacade(mockRepo)

	category := &domain.Category{ID: 1, Name: "Test Category", Type: "Expense"}

	mockRepo.On("SaveCategory", category).Return(nil)

	err := facade.CreateCategory(category)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCategoryFacade_GetCategory(t *testing.T) {
	mockRepo := mocks.NewRepository(t)
	facade := facades.NewCategoryFacade(mockRepo)

	category := &domain.Category{ID: 1, Name: "Test Category", Type: "Expense"}

	mockRepo.On("GetCategory", 1).Return(category, nil)

	retrievedCategory, err := facade.GetCategory(1)
	assert.NoError(t, err)
	assert.Equal(t, category, retrievedCategory)
	mockRepo.AssertExpectations(t)
}

func TestCategoryFacade_DeleteCategory(t *testing.T) {
	mockRepo := mocks.NewRepository(t)
	facade := facades.NewCategoryFacade(mockRepo)

	mockRepo.On("DeleteCategory", 1).Return(nil)

	err := facade.DeleteCategory(1)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCategoryFacade_UpdateCategory(t *testing.T) {
	mockRepo := mocks.NewRepository(t)
	facade := facades.NewCategoryFacade(mockRepo)

	category := &domain.Category{ID: 1, Name: "Updated Category", Type: "Income"}

	mockRepo.On("UpdateCategory", category).Return(nil)

	err := facade.UpdateCategory(category)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCategoryFacade_GetCategory_NotFound(t *testing.T) {
	mockRepo := mocks.NewRepository(t)
	facade := facades.NewCategoryFacade(mockRepo)

	mockRepo.On("GetCategory", 1).Return(nil, errors.New("category not found"))

	_, err := facade.GetCategory(1)
	assert.Error(t, err)
	assert.EqualError(t, err, "category not found")
	mockRepo.AssertExpectations(t)
}
