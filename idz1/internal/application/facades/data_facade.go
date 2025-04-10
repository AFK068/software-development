package facades

import (
	"fmt"

	"github.com/AFK068/bot/internal/domain"
	"github.com/AFK068/bot/internal/infrastructure/exporter"
	"github.com/AFK068/bot/internal/infrastructure/importer"
)

type DataService interface {
	ExportCategories(visitor exporter.Exporter, fileName string) error
	ExportBankAccounts(visitor exporter.Exporter, fileName string) error
	ExportOperations(visitor exporter.Exporter, fileName string) error
	ImportCategories(visitor importer.Importer, fileName string) error
	ImportBankAccounts(visitor importer.Importer, fileName string) error
	ImportOperations(visitor importer.Importer, fileName string) error
}

type DataFacade struct {
	repo domain.Repository
}

func NewDataFacade(repo domain.Repository) *DataFacade {
	return &DataFacade{repo: repo}
}

func (f *DataFacade) ExportCategories(visitor exporter.Exporter, fileName string) error {
	categories, err := f.repo.GetAllCategories()
	if err != nil {
		return fmt.Errorf("failed to get categories: %w", err)
	}

	for _, category := range categories {
		if err := visitor.VisitCategory(category); err != nil {
			return fmt.Errorf("failed to visit category: %w", err)
		}
	}

	return visitor.Export(fileName)
}

func (f *DataFacade) ExportBankAccounts(visitor exporter.Exporter, fileName string) error {
	accounts, err := f.repo.GetAllBankAccounts()
	if err != nil {
		return fmt.Errorf("failed to get bank accounts: %w", err)
	}

	for _, account := range accounts {
		if err := visitor.VisitBankAccount(account); err != nil {
			return fmt.Errorf("failed to visit bank account: %w", err)
		}
	}

	return visitor.Export(fileName)
}

func (f *DataFacade) ExportOperations(visitor exporter.Exporter, fileName string) error {
	operations, err := f.repo.GetAllOperations()
	if err != nil {
		return fmt.Errorf("failed to get operations: %w", err)
	}

	for _, operation := range operations {
		if err := visitor.VisitOperation(operation); err != nil {
			return fmt.Errorf("failed to visit operation: %w", err)
		}
	}

	return visitor.Export(fileName)
}

func (f *DataFacade) ImportCategories(visitor importer.Importer, fileName string) error {
	categories, err := visitor.Import(fileName)
	if err != nil {
		return fmt.Errorf("failed to import categories: %w", err)
	}

	for _, category := range categories.([]*domain.Category) {
		if err := f.repo.SaveCategory(category); err != nil {
			fmt.Printf("failed to save category: %v\n", err)
		}
	}

	return nil
}

func (f *DataFacade) ImportBankAccounts(visitor importer.Importer, fileName string) error {
	accounts, err := visitor.Import(fileName)
	if err != nil {
		return fmt.Errorf("failed to import bank accounts: %w", err)
	}

	for _, account := range accounts.([]*domain.BankAccount) {
		if err := f.repo.SaveBankAccount(account); err != nil {
			fmt.Printf("failed to save bank account: %v\n", err)
		}
	}

	return nil
}

func (f *DataFacade) ImportOperations(visitor importer.Importer, fileName string) error {
	operations, err := visitor.Import(fileName)
	if err != nil {
		return fmt.Errorf("failed to import operations: %w", err)
	}

	for _, operation := range operations.([]*domain.Operation) {
		if err := f.repo.SaveOperation(operation); err != nil {
			fmt.Printf("failed to save operation: %v\n", err)
		}
	}

	return nil
}
