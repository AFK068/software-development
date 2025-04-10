package repository_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/AFK068/bot/internal/domain"
	"github.com/AFK068/bot/internal/infrastructure/repository"
)

func TestSaveAndGetBankAccount(t *testing.T) {
	repo := repository.NewIInMemoryRepository()
	account := &domain.BankAccount{ID: 1, Name: "Test Account", Balance: 1000}
	err := repo.SaveBankAccount(account)
	assert.NoError(t, err)

	retrievedAccount, err := repo.GetBankAccount(1)
	assert.NoError(t, err)
	assert.Equal(t, account, retrievedAccount)
}

func TestUpdateBankAccount(t *testing.T) {
	repo := repository.NewIInMemoryRepository()
	account := &domain.BankAccount{ID: 1, Name: "Test Account", Balance: 1000}
	_ = repo.SaveBankAccount(account)

	account.Balance = 2000
	err := repo.UpdateBankAccount(account)
	assert.NoError(t, err)

	updatedAccount, err := repo.GetBankAccount(1)
	assert.NoError(t, err)
	assert.Equal(t, 2000.0, updatedAccount.Balance)
}

func TestDeleteBankAccount(t *testing.T) {
	repo := repository.NewIInMemoryRepository()
	account := &domain.BankAccount{ID: 1, Name: "Test Account", Balance: 1000}
	_ = repo.SaveBankAccount(account)

	err := repo.DeleteBankAccount(1)
	assert.NoError(t, err)

	_, err = repo.GetBankAccount(1)
	assert.Error(t, err)
}

func TestSaveAndGetCategory(t *testing.T) {
	repo := repository.NewIInMemoryRepository()
	category := &domain.Category{ID: 1, Name: "Test Category", Type: "Expense"}
	err := repo.SaveCategory(category)
	assert.NoError(t, err)

	retrievedCategory, err := repo.GetCategory(1)
	assert.NoError(t, err)
	assert.Equal(t, category, retrievedCategory)
}

func TestUpdateCategory(t *testing.T) {
	repo := repository.NewIInMemoryRepository()
	category := &domain.Category{ID: 1, Name: "Test Category", Type: "Expense"}
	_ = repo.SaveCategory(category)

	category.Name = "Updated Category"
	err := repo.UpdateCategory(category)
	assert.NoError(t, err)

	updatedCategory, err := repo.GetCategory(1)
	assert.NoError(t, err)
	assert.Equal(t, "Updated Category", updatedCategory.Name)
}

func TestDeleteCategory(t *testing.T) {
	repo := repository.NewIInMemoryRepository()
	category := &domain.Category{ID: 1, Name: "Test Category", Type: "Expense"}
	_ = repo.SaveCategory(category)

	err := repo.DeleteCategory(1)
	assert.NoError(t, err)

	_, err = repo.GetCategory(1)
	assert.Error(t, err)
}
