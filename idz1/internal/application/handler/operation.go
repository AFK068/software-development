package handler

import (
	"errors"
	"math"

	"github.com/AFK068/bot/internal/application/facades"
	"github.com/AFK068/bot/internal/domain"
	"github.com/AFK068/bot/internal/infrastructure/userinteraction"
)

type OperationHandler interface {
	CreateOperation() error
	UpdateOperation() error
	DeleteOperation() error
}

type DefaultOperationHandler struct {
	operationFacade facades.OperationService
	userInput       userinteraction.UserInput
}

func NewOperationHandler(facade facades.OperationService, userInput userinteraction.UserInput) *DefaultOperationHandler {
	return &DefaultOperationHandler{
		operationFacade: facade,
		userInput:       userInput,
	}
}

func (h *DefaultOperationHandler) CreateOperation() error {
	operation, err := h.collectOperationData(nil)
	if err != nil {
		return err
	}

	err = h.operationFacade.CreateOperation(operation)
	if err != nil {
		return err
	}

	return nil
}

func (h *DefaultOperationHandler) UpdateOperation() error {
	idValue, err := h.userInput.GetIntValue("Enter the ID of the operation", 0, math.MaxInt)
	if err != nil {
		return err
	}

	operation, err := h.operationFacade.GetOperation(idValue)
	if err != nil {
		return err
	}

	updatedOperation, err := h.collectOperationData(operation)
	if err != nil {
		return err
	}

	err = h.operationFacade.UpdateOperation(updatedOperation)
	if err != nil {
		return err
	}

	return nil
}

func (h *DefaultOperationHandler) DeleteOperation() error {
	idValue, err := h.userInput.GetIntValue("Enter the ID of the operation", 0, math.MaxInt)
	if err != nil {
		return err
	}

	return h.operationFacade.DeleteOperation(idValue)
}

func (h *DefaultOperationHandler) collectOperationData(existingOperation *domain.Operation) (*domain.Operation, error) {
	var operation domain.Operation
	if existingOperation != nil {
		operation = *existingOperation
	}

	idValue := operation.ID

	if existingOperation == nil {
		var err error

		idValue, err = h.userInput.GetIntValue("Enter the ID of the operation", 0, math.MaxInt)
		if err != nil {
			return nil, err
		}
	}

	typeValue, err := h.userInput.GetStringValue("Enter the type of the operation (income/expense)")
	if err != nil {
		return nil, err
	}

	if typeValue != domain.OperationTypeIncome && typeValue != domain.OperationTypeExpense {
		return nil, errors.New("invalid operation type")
	}

	bankAccountID, err := h.userInput.GetIntValue("Enter the bank account ID", 0, math.MaxInt)
	if err != nil {
		return nil, err
	}

	amount, err := h.userInput.GetFloatValue("Enter the amount of the operation", 0, math.MaxFloat64)
	if err != nil {
		return nil, err
	}

	date, err := h.userInput.GetTimeValue("Enter the date of the operation (YYYY-MM-DD HH:MM:SS)")
	if err != nil {
		return nil, err
	}

	description, err := h.userInput.GetStringValue("Enter a description for the operation")
	if err != nil {
		return nil, err
	}

	categoryID, err := h.userInput.GetIntValue("Enter the category ID", 0, math.MaxInt)
	if err != nil {
		return nil, err
	}

	return &domain.Operation{
		ID:            idValue,
		Type:          typeValue,
		BankAccountID: bankAccountID,
		Amount:        amount,
		Date:          date,
		Description:   description,
		CategoryID:    categoryID,
	}, nil
}
