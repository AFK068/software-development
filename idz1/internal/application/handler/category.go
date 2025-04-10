package handler

import (
	"math"

	"github.com/AFK068/bot/internal/application/facades"
	"github.com/AFK068/bot/internal/domain"
	"github.com/AFK068/bot/internal/infrastructure/userinteraction"
)

type CategoryHandler interface {
	CreateCategory() error
	UpdateCategory() error
	DeleteCategory() error
}

type DefaultCategoryHandler struct {
	categoryFacade facades.CategoryService
	userInput      userinteraction.UserInput
}

func NewCategoryHandler(facade facades.CategoryService, userInput userinteraction.UserInput) *DefaultCategoryHandler {
	return &DefaultCategoryHandler{
		categoryFacade: facade,
		userInput:      userInput,
	}
}

func (h *DefaultCategoryHandler) CreateCategory() error {
	category, err := h.collectCategoryData(nil)
	if err != nil {
		return err
	}

	err = h.categoryFacade.CreateCategory(category)
	if err != nil {
		return err
	}

	return nil
}

func (h *DefaultCategoryHandler) UpdateCategory() error {
	idValue, err := h.userInput.GetIntValue("Enter the category ID", 0, math.MaxInt)
	if err != nil {
		return err
	}

	category, err := h.categoryFacade.GetCategory(idValue)
	if err != nil {
		return err
	}

	updatedCategory, err := h.collectCategoryData(category)
	if err != nil {
		return err
	}

	err = h.categoryFacade.UpdateCategory(updatedCategory)
	if err != nil {
		return err
	}

	return nil
}

func (h *DefaultCategoryHandler) DeleteCategory() error {
	idValue, err := h.userInput.GetIntValue("Enter the category ID to delete", 0, math.MaxInt)
	if err != nil {
		return err
	}

	return h.categoryFacade.DeleteCategory(idValue)
}

func (h *DefaultCategoryHandler) collectCategoryData(existingCategory *domain.Category) (*domain.Category, error) {
	var category domain.Category
	if existingCategory != nil {
		category = *existingCategory
	}

	idValue := category.ID

	if existingCategory == nil {
		var err error

		idValue, err = h.userInput.GetIntValue("Enter the category ID", 0, math.MaxInt)
		if err != nil {
			return nil, err
		}
	}

	name, err := h.userInput.GetStringValue("Enter the category name")
	if err != nil {
		return nil, err
	}

	categoryType, err := h.userInput.GetStringValue("Enter the category type (e.g., income/expense)")
	if err != nil {
		return nil, err
	}

	return &domain.Category{
		ID:   idValue,
		Name: name,
		Type: categoryType,
	}, nil
}
