package repository

import (
	"errors"

	"github.com/AFK068/bot/internal/domain"
)

type InMemoryRepo struct {
	bankAccounts map[int]*domain.BankAccount
	categories   map[int]*domain.Category
	operations   map[int]*domain.Operation
}

func NewIInMemoryRepository() *InMemoryRepo {
	return &InMemoryRepo{
		bankAccounts: make(map[int]*domain.BankAccount),
		categories:   make(map[int]*domain.Category),
		operations:   make(map[int]*domain.Operation),
	}
}

func (r *InMemoryRepo) SaveBankAccount(bankAccount *domain.BankAccount) error {
	r.bankAccounts[bankAccount.ID] = bankAccount
	return nil
}

func (r *InMemoryRepo) GetBankAccount(id int) (*domain.BankAccount, error) {
	bankAccount, ok := r.bankAccounts[id]
	if !ok {
		return nil, errors.New("bank account not found")
	}

	return bankAccount, nil
}

func (r *InMemoryRepo) DeleteBankAccount(id int) error {
	if _, ok := r.bankAccounts[id]; !ok {
		return errors.New("bank account not found")
	}

	delete(r.bankAccounts, id)

	return nil
}

func (r *InMemoryRepo) UpdateBankAccount(bankAccount *domain.BankAccount) error {
	if _, ok := r.bankAccounts[bankAccount.ID]; !ok {
		return errors.New("bank account not found")
	}

	r.bankAccounts[bankAccount.ID] = bankAccount

	return nil
}

func (r *InMemoryRepo) GetAllBankAccounts() ([]*domain.BankAccount, error) {
	accounts := make([]*domain.BankAccount, 0, len(r.bankAccounts))
	for _, account := range r.bankAccounts {
		accounts = append(accounts, account)
	}

	return accounts, nil
}

func (r *InMemoryRepo) SaveCategory(category *domain.Category) error {
	r.categories[category.ID] = category
	return nil
}

func (r *InMemoryRepo) GetCategory(id int) (*domain.Category, error) {
	category, ok := r.categories[id]
	if !ok {
		return nil, errors.New("category not found")
	}

	return category, nil
}

func (r *InMemoryRepo) DeleteCategory(id int) error {
	if _, ok := r.categories[id]; !ok {
		return errors.New("category not found")
	}

	delete(r.categories, id)

	return nil
}

func (r *InMemoryRepo) UpdateCategory(category *domain.Category) error {
	if _, ok := r.categories[category.ID]; !ok {
		return errors.New("category not found")
	}

	r.categories[category.ID] = category

	return nil
}

func (r *InMemoryRepo) GetAllCategories() ([]*domain.Category, error) {
	categories := make([]*domain.Category, 0, len(r.categories))
	for _, category := range r.categories {
		categories = append(categories, category)
	}

	return categories, nil
}

func (r *InMemoryRepo) SaveOperation(operation *domain.Operation) error {
	r.operations[operation.ID] = operation
	return nil
}

func (r *InMemoryRepo) GetOperation(id int) (*domain.Operation, error) {
	operation, ok := r.operations[id]
	if !ok {
		return nil, errors.New("operation not found")
	}

	return operation, nil
}

func (r *InMemoryRepo) DeleteOperation(id int) error {
	if _, ok := r.operations[id]; !ok {
		return errors.New("operation not found")
	}

	delete(r.operations, id)

	return nil
}

func (r *InMemoryRepo) UpdateOperation(operation *domain.Operation) error {
	if _, ok := r.operations[operation.ID]; !ok {
		return errors.New("operation not found")
	}

	r.operations[operation.ID] = operation

	return nil
}

func (r *InMemoryRepo) GetAllOperations() ([]*domain.Operation, error) {
	operations := make([]*domain.Operation, 0, len(r.operations))
	for _, operation := range r.operations {
		operations = append(operations, operation)
	}

	return operations, nil
}

func (r *InMemoryRepo) GetOperationsByAccountID(accountID int) ([]*domain.Operation, error) {
	operations := make([]*domain.Operation, 0)
	for _, operation := range r.operations {
		if operation.ID == accountID {
			operations = append(operations, operation)
		}
	}

	return operations, nil
}
