package factories

import (
	"errors"
	"time"

	"github.com/AFK068/bot/internal/domain"
)

type BankFactory struct{}

func NewBankFactory() *BankFactory {
	return &BankFactory{}
}

func (b *BankFactory) NewBankAccount(id int, name string, balance float64) (*domain.BankAccount, error) {
	if name == "" {
		return nil, errors.New("name is required")
	}

	if balance < 0 {
		return nil, errors.New("negative balance")
	}

	return &domain.BankAccount{
		ID:      id,
		Name:    name,
		Balance: balance,
	}, nil
}

func (b *BankFactory) NewCategory(id int, name, categoryType string) (*domain.Category, error) {
	if name == "" {
		return nil, errors.New("name is required")
	}

	if categoryType == "" {
		return nil, errors.New("category type is required")
	}

	return &domain.Category{
		ID:   id,
		Name: name,
		Type: categoryType,
	}, nil
}

func (b *BankFactory) NewOperation(id int, typeOperation string,
	bankAccountID int, amount float64, date time.Time, description string, categoryID int) (*domain.Operation, error) {
	if typeOperation == "" {
		return nil, errors.New("type is required")
	}

	if amount <= 0 {
		return nil, errors.New("negative amount")
	}

	if date.IsZero() {
		return nil, errors.New("date is required")
	}

	if description == "" {
		return nil, errors.New("description is required")
	}

	return &domain.Operation{
		ID:            id,
		Type:          typeOperation,
		BankAccountID: bankAccountID,
		Amount:        amount,
		Date:          date,
		Description:   description,
		CategoryID:    categoryID,
	}, nil
}
