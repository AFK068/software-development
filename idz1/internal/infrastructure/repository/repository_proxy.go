package repository

import (
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"

	"github.com/AFK068/bot/internal/domain"
)

const (
	DefaultExpiration = 5 * time.Minute
	DefaultCleanup    = 10 * time.Minute
)

type ProxyRepo struct {
	repo  domain.Repository
	cache *cache.Cache
}

func NewProxy(repo domain.Repository, expiration, cleanup time.Duration) *ProxyRepo {
	return &ProxyRepo{
		repo:  repo,
		cache: cache.New(expiration, cleanup),
	}
}

func (p *ProxyRepo) SaveBankAccount(bankAccount *domain.BankAccount) error {
	key := fmt.Sprintf("bank_account_%d", bankAccount.ID)
	p.cache.Set(key, bankAccount, cache.DefaultExpiration)

	return p.repo.SaveBankAccount(bankAccount)
}

func (p *ProxyRepo) GetBankAccount(id int) (*domain.BankAccount, error) {
	key := fmt.Sprintf("bank_account_%d", id)
	if bankAccount, found := p.cache.Get(key); found {
		if bankAccount, ok := bankAccount.(*domain.BankAccount); ok {
			return bankAccount, nil
		}
	}

	account, err := p.repo.GetBankAccount(id)
	if err != nil {
		return nil, err
	}

	p.cache.Set(key, account, cache.DefaultExpiration)

	return account, nil
}

func (p *ProxyRepo) DeleteBankAccount(id int) error {
	key := fmt.Sprintf("bank_account_%d", id)
	if _, found := p.cache.Get(key); found {
		p.cache.Delete(key)
	}

	return p.repo.DeleteBankAccount(id)
}

func (p *ProxyRepo) UpdateBankAccount(bankAccount *domain.BankAccount) error {
	key := fmt.Sprintf("bank_account_%d", bankAccount.ID)
	p.cache.Set(key, bankAccount, cache.DefaultExpiration)

	return p.repo.UpdateBankAccount(bankAccount)
}

func (p *ProxyRepo) GetAllBankAccounts() ([]*domain.BankAccount, error) {
	accounts, err := p.repo.GetAllBankAccounts()
	if err != nil {
		return nil, err
	}

	for _, account := range accounts {
		key := fmt.Sprintf("bank_account_%d", account.ID)
		p.cache.Set(key, account, cache.DefaultExpiration)
	}

	return accounts, nil
}

func (p *ProxyRepo) SaveCategory(category *domain.Category) error {
	key := fmt.Sprintf("category_%d", category.ID)
	p.cache.Set(key, category, cache.DefaultExpiration)

	return p.repo.SaveCategory(category)
}

func (p *ProxyRepo) GetCategory(id int) (*domain.Category, error) {
	key := fmt.Sprintf("category_%d", id)
	if category, found := p.cache.Get(key); found {
		if category, ok := category.(*domain.Category); ok {
			return category, nil
		}
	}

	category, err := p.repo.GetCategory(id)
	if err != nil {
		return nil, err
	}

	p.cache.Set(key, category, cache.DefaultExpiration)

	return category, nil
}

func (p *ProxyRepo) DeleteCategory(id int) error {
	key := fmt.Sprintf("category_%d", id)
	if _, found := p.cache.Get(key); found {
		p.cache.Delete(key)
	}

	return p.repo.DeleteCategory(id)
}

func (p *ProxyRepo) UpdateCategory(category *domain.Category) error {
	key := fmt.Sprintf("category_%d", category.ID)
	p.cache.Set(key, category, cache.DefaultExpiration)

	return p.repo.UpdateCategory(category)
}

func (p *ProxyRepo) GetAllCategories() ([]*domain.Category, error) {
	categories, err := p.repo.GetAllCategories()
	if err != nil {
		return nil, err
	}

	for _, category := range categories {
		key := fmt.Sprintf("category_%d", category.ID)
		p.cache.Set(key, category, cache.DefaultExpiration)
	}

	return categories, nil
}

func (p *ProxyRepo) SaveOperation(operation *domain.Operation) error {
	key := fmt.Sprintf("operation_%d", operation.ID)
	p.cache.Set(key, operation, cache.DefaultExpiration)

	return p.repo.SaveOperation(operation)
}

func (p *ProxyRepo) GetOperation(id int) (*domain.Operation, error) {
	key := fmt.Sprintf("operation_%d", id)
	if operation, found := p.cache.Get(key); found {
		if operation, ok := operation.(*domain.Operation); ok {
			return operation, nil
		}
	}

	operation, err := p.repo.GetOperation(id)
	if err != nil {
		return nil, err
	}

	p.cache.Set(key, operation, cache.DefaultExpiration)

	return operation, nil
}

func (p *ProxyRepo) DeleteOperation(id int) error {
	key := fmt.Sprintf("operation_%d", id)
	if _, found := p.cache.Get(key); found {
		p.cache.Delete(key)
	}

	return p.repo.DeleteOperation(id)
}

func (p *ProxyRepo) UpdateOperation(operation *domain.Operation) error {
	key := fmt.Sprintf("operation_%d", operation.ID)
	p.cache.Set(key, operation, cache.DefaultExpiration)

	return p.repo.UpdateOperation(operation)
}

func (p *ProxyRepo) GetAllOperations() ([]*domain.Operation, error) {
	operations, err := p.repo.GetAllOperations()
	if err != nil {
		return nil, err
	}

	for _, operation := range operations {
		key := fmt.Sprintf("operation_%d", operation.ID)
		p.cache.Set(key, operation, cache.DefaultExpiration)
	}

	return operations, nil
}

func (p *ProxyRepo) GetOperationsByAccountID(accountID int) ([]*domain.Operation, error) {
	key := fmt.Sprintf("operations_%d", accountID)
	if operations, found := p.cache.Get(key); found {
		if operations, ok := operations.([]*domain.Operation); ok {
			return operations, nil
		}
	}

	operations, err := p.repo.GetOperationsByAccountID(accountID)
	if err != nil {
		return nil, err
	}

	p.cache.Set(key, operations, cache.DefaultExpiration)

	return operations, nil
}
