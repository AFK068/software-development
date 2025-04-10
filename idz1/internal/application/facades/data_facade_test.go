package facades_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/AFK068/bot/internal/application/facades"
	"github.com/AFK068/bot/internal/domain"

	repoMock "github.com/AFK068/bot/internal/domain/mocks"
	exportMock "github.com/AFK068/bot/internal/infrastructure/exporter/mocks"
	importMock "github.com/AFK068/bot/internal/infrastructure/importer/mocks"
)

func TestDataFacade_ExportCategories(t *testing.T) {
	mockRepo := repoMock.NewRepository(t)
	mockExporter := exportMock.NewExporter(t)
	facade := facades.NewDataFacade(mockRepo)

	categories := []*domain.Category{
		{ID: 1, Name: "Test Category", Type: "Expense"},
	}

	mockRepo.On("GetAllCategories").Return(categories, nil)
	mockExporter.On("VisitCategory", categories[0]).Return(nil)
	mockExporter.On("Export", "categories.yaml").Return(nil)

	err := facade.ExportCategories(mockExporter, "categories.yaml")
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
	mockExporter.AssertExpectations(t)
}

func TestDataFacade_ImportCategories(t *testing.T) {
	mockRepo := repoMock.NewRepository(t)
	mockImporter := importMock.NewImporter(t)
	facade := facades.NewDataFacade(mockRepo)

	categories := []*domain.Category{
		{ID: 1, Name: "Test Category", Type: "Expense"},
	}

	mockImporter.On("Import", "categories.yaml").Return(categories, nil)
	mockRepo.On("SaveCategory", categories[0]).Return(nil)

	err := facade.ImportCategories(mockImporter, "categories.yaml")
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
	mockImporter.AssertExpectations(t)
}

func TestDataFacade_ExportBankAccounts(t *testing.T) {
	mockRepo := repoMock.NewRepository(t)
	mockExporter := exportMock.NewExporter(t)
	facade := facades.NewDataFacade(mockRepo)

	accounts := []*domain.BankAccount{
		{ID: 1, Name: "Test Account", Balance: 1000},
	}

	mockRepo.On("GetAllBankAccounts").Return(accounts, nil)
	mockExporter.On("VisitBankAccount", accounts[0]).Return(nil)
	mockExporter.On("Export", "accounts.yaml").Return(nil)

	err := facade.ExportBankAccounts(mockExporter, "accounts.yaml")
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
	mockExporter.AssertExpectations(t)
}

func TestDataFacade_ImportBankAccounts(t *testing.T) {
	mockRepo := repoMock.NewRepository(t)
	mockImporter := importMock.NewImporter(t)
	facade := facades.NewDataFacade(mockRepo)

	accounts := []*domain.BankAccount{
		{ID: 1, Name: "Test Account", Balance: 1000},
	}

	mockImporter.On("Import", "accounts.yaml").Return(accounts, nil)
	mockRepo.On("SaveBankAccount", accounts[0]).Return(nil)

	err := facade.ImportBankAccounts(mockImporter, "accounts.yaml")
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
	mockImporter.AssertExpectations(t)
}

func TestDataFacade_ExportOperations(t *testing.T) {
	mockRepo := repoMock.NewRepository(t)
	mockExporter := exportMock.NewExporter(t)
	facade := facades.NewDataFacade(mockRepo)

	operations := []*domain.Operation{
		{ID: 1, Type: domain.OperationTypeIncome, Amount: 1000},
	}

	mockRepo.On("GetAllOperations").Return(operations, nil)
	mockExporter.On("VisitOperation", operations[0]).Return(nil)
	mockExporter.On("Export", "operations.yaml").Return(nil)

	err := facade.ExportOperations(mockExporter, "operations.yaml")
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
	mockExporter.AssertExpectations(t)
}

func TestDataFacade_ImportOperations(t *testing.T) {
	mockRepo := repoMock.NewRepository(t)
	mockImporter := importMock.NewImporter(t)
	facade := facades.NewDataFacade(mockRepo)

	operations := []*domain.Operation{
		{ID: 1, Type: domain.OperationTypeIncome, Amount: 1000},
	}

	mockImporter.On("Import", "operations.yaml").Return(operations, nil)
	mockRepo.On("SaveOperation", operations[0]).Return(nil)

	err := facade.ImportOperations(mockImporter, "operations.yaml")
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
	mockImporter.AssertExpectations(t)
}
