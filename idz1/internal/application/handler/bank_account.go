package handler

import (
	"math"

	"github.com/AFK068/bot/internal/application/facades"
	"github.com/AFK068/bot/internal/domain"
	"github.com/AFK068/bot/internal/infrastructure/userinteraction"
)

type BankAccountHandler interface {
	CreateBankAccount() error
	UpdateBankAccount() error
	DeleteBankAccount() error
}

type DefaultBankAccountHandler struct {
	accountFacade facades.BankAccountService
	userInput     userinteraction.UserInput
}

func NewBankAccountHandler(facade facades.BankAccountService, userInput userinteraction.UserInput) *DefaultBankAccountHandler {
	return &DefaultBankAccountHandler{
		accountFacade: facade,
		userInput:     userInput,
	}
}

func (h *DefaultBankAccountHandler) CreateBankAccount() error {
	account, err := h.collectBankAccountData(nil)
	if err != nil {
		return err
	}

	err = h.accountFacade.CreateBankAccount(account)
	if err != nil {
		return err
	}

	return nil
}

func (h *DefaultBankAccountHandler) UpdateBankAccount() error {
	idValue, err := h.userInput.GetIntValue("Enter the account ID", 0, math.MaxInt)
	if err != nil {
		return err
	}

	account, err := h.accountFacade.GetBankAccount(idValue)
	if err != nil {
		return err
	}

	updatedAccount, err := h.collectBankAccountData(account)
	if err != nil {
		return err
	}

	err = h.accountFacade.UpdateBankAccount(updatedAccount)
	if err != nil {
		return err
	}

	return nil
}

func (h *DefaultBankAccountHandler) DeleteBankAccount() error {
	idValue, err := h.userInput.GetIntValue("Enter the account ID to delete", 0, math.MaxInt)
	if err != nil {
		return err
	}

	return h.accountFacade.DeleteBankAccount(idValue)
}

func (h *DefaultBankAccountHandler) collectBankAccountData(existingAccount *domain.BankAccount) (*domain.BankAccount, error) {
	var account domain.BankAccount
	if existingAccount != nil {
		account = *existingAccount
	}

	idValue := account.ID

	if existingAccount == nil {
		var err error

		idValue, err = h.userInput.GetIntValue("Enter the account ID", 0, math.MaxInt)
		if err != nil {
			return nil, err
		}
	}

	name, err := h.userInput.GetStringValue("Enter the account name")
	if err != nil {
		return nil, err
	}

	balance, err := h.userInput.GetFloatValue("Enter the account balance", 0, math.MaxFloat64)
	if err != nil {
		return nil, err
	}

	return &domain.BankAccount{
		ID:      idValue,
		Name:    name,
		Balance: balance,
	}, nil
}
