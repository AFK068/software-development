package facades_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/AFK068/bot/internal/application/facades"
	"github.com/AFK068/bot/internal/domain"
	"github.com/AFK068/bot/internal/domain/mocks"
)

func TestBankAccountFacade_CreateAccount(t *testing.T) {
	mockRepo := mocks.NewRepository(t)
	facade := facades.NewBankAccountFacade(mockRepo)

	account := &domain.BankAccount{ID: 1, Name: "Test Account", Balance: 1000}

	mockRepo.On("SaveBankAccount", account).Return(nil)

	err := facade.CreateBankAccount(account)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestBankAccountFacade_GetAccount(t *testing.T) {
	mockRepo := mocks.NewRepository(t)
	facade := facades.NewBankAccountFacade(mockRepo)

	account := &domain.BankAccount{ID: 1, Name: "Test Account", Balance: 1000}

	mockRepo.On("GetBankAccount", 1).Return(account, nil)

	retrievedAccount, err := facade.GetBankAccount(1)
	assert.NoError(t, err)
	assert.Equal(t, account, retrievedAccount)
	mockRepo.AssertExpectations(t)
}

func TestBankAccountFacade_DeleteAccount(t *testing.T) {
	mockRepo := mocks.NewRepository(t)
	facade := facades.NewBankAccountFacade(mockRepo)

	mockRepo.On("DeleteBankAccount", 1).Return(nil)

	err := facade.DeleteBankAccount(1)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestBankAccountFacade_UpdateAccount(t *testing.T) {
	mockRepo := mocks.NewRepository(t)
	facade := facades.NewBankAccountFacade(mockRepo)

	account := &domain.BankAccount{ID: 1, Name: "Updated Account", Balance: 2000}

	mockRepo.On("UpdateBankAccount", account).Return(nil)

	err := facade.UpdateBankAccount(account)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestBankAccountFacade_GetAccount_NotFound(t *testing.T) {
	mockRepo := mocks.NewRepository(t)
	facade := facades.NewBankAccountFacade(mockRepo)

	mockRepo.On("GetBankAccount", 1).Return(nil, errors.New("account not found"))

	_, err := facade.GetBankAccount(1)
	assert.Error(t, err)
	assert.EqualError(t, err, "account not found")
	mockRepo.AssertExpectations(t)
}
