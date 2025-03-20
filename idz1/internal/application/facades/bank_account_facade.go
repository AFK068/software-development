package facades

import (
	"github.com/AFK068/bot/internal/domain"
)

type BankAccountService interface {
	CreateBankAccount(account *domain.BankAccount) error
	GetBankAccount(id int) (*domain.BankAccount, error)
	DeleteBankAccount(id int) error
	UpdateBankAccount(account *domain.BankAccount) error
}

type BankAccountFacade struct {
	repo domain.Repository
}

func NewBankAccountFacade(repo domain.Repository) *BankAccountFacade {
	return &BankAccountFacade{repo: repo}
}

func (f *BankAccountFacade) CreateBankAccount(account *domain.BankAccount) error {
	return f.repo.SaveBankAccount(account)
}

func (f *BankAccountFacade) GetBankAccount(id int) (*domain.BankAccount, error) {
	return f.repo.GetBankAccount(id)
}

func (f *BankAccountFacade) DeleteBankAccount(id int) error {
	return f.repo.DeleteBankAccount(id)
}

func (f *BankAccountFacade) UpdateBankAccount(account *domain.BankAccount) error {
	return f.repo.UpdateBankAccount(account)
}
