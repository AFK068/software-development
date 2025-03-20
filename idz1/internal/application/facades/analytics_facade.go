package facades

import "github.com/AFK068/bot/internal/domain"

type AnalyticsService interface {
	CalculateBalanceDifference() (float64, error)
	CalculateBalance() (float64, error)
	RecalculateBalances() error
}

type AnalyticsFacade struct {
	repo domain.Repository
}

func NewAnalyticsFacade(repo domain.Repository) *AnalyticsFacade {
	return &AnalyticsFacade{repo: repo}
}

func (f *AnalyticsFacade) CalculateBalanceDifference() (float64, error) {
	var diff float64

	operations, err := f.repo.GetAllOperations()
	if err != nil {
		return 0, err
	}

	for _, op := range operations {
		switch op.Type {
		case domain.OperationTypeIncome:
			diff += op.Amount
		case domain.OperationTypeExpense:
			diff -= op.Amount
		}
	}

	return diff, nil
}

func (f *AnalyticsFacade) CalculateBalance() (float64, error) {
	var balance float64

	accounts, err := f.repo.GetAllBankAccounts()
	if err != nil {
		return 0, err
	}

	for _, acc := range accounts {
		balance += acc.Balance
	}

	return balance, nil
}

func (f *AnalyticsFacade) RecalculateBalances() error {
	accounts, err := f.repo.GetAllBankAccounts()
	if err != nil {
		return err
	}

	for _, acc := range accounts {
		operations, err := f.repo.GetOperationsByAccountID(acc.ID)
		if err != nil {
			return err
		}

		for _, op := range operations {
			switch op.Type {
			case domain.OperationTypeIncome:
				acc.Balance += op.Amount
			case domain.OperationTypeExpense:
				acc.Balance -= op.Amount
			}
		}

		if err := f.repo.SaveBankAccount(acc); err != nil {
			return err
		}
	}

	return nil
}
