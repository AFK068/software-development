package facades

import "github.com/AFK068/bot/internal/domain"

type OperationService interface {
	CreateOperation(op *domain.Operation) error
	GetOperation(id int) (*domain.Operation, error)
	DeleteOperation(id int) error
	UpdateOperation(op *domain.Operation) error
}

type OperationFacade struct {
	repo domain.Repository
}

func NewOperationFacade(repo domain.Repository) *OperationFacade {
	return &OperationFacade{repo: repo}
}

func (f *OperationFacade) CreateOperation(op *domain.Operation) error {
	account, err := f.repo.GetBankAccount(op.ID)
	if err != nil {
		return err
	}

	switch op.Type {
	case domain.OperationTypeIncome:
		account.Balance += op.Amount
	case domain.OperationTypeExpense:
		account.Balance -= op.Amount
	}

	if err := f.repo.SaveBankAccount(account); err != nil {
		return err
	}

	return f.repo.SaveOperation(op)
}

func (f *OperationFacade) GetOperation(id int) (*domain.Operation, error) {
	return f.repo.GetOperation(id)
}

func (f *OperationFacade) DeleteOperation(id int) error {
	return f.repo.DeleteOperation(id)
}

func (f *OperationFacade) UpdateOperation(op *domain.Operation) error {
	return f.repo.UpdateOperation(op)
}
