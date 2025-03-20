package main

import (
	"fmt"
	"log"
	"os"

	"github.com/manifoldco/promptui"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"

	"github.com/AFK068/bot/internal/application/facades"
	"github.com/AFK068/bot/internal/application/handler"
	"github.com/AFK068/bot/internal/domain"
	"github.com/AFK068/bot/internal/domain/factories"
	"github.com/AFK068/bot/internal/infrastructure/repository"
	"github.com/AFK068/bot/internal/infrastructure/userinteraction"
)

func main() {
	fx.New(
		fx.Provide(
			// Repo.
			repository.NewIInMemoryRepository,
			repository.NewProxy,

			func(inMemoryRepo *repository.InMemoryRepo) domain.Repository {
				return repository.NewProxy(inMemoryRepo, repository.DefaultExpiration, repository.DefaultCleanup)
			},

			// Facades.
			fx.Annotate(
				facades.NewBankAccountFacade,
				fx.As(new(facades.BankAccountService)),
			),

			fx.Annotate(
				facades.NewCategoryFacade,
				fx.As(new(facades.CategoryService)),
			),

			fx.Annotate(
				facades.NewOperationFacade,
				fx.As(new(facades.OperationService)),
			),

			facades.NewAnalyticsFacade,

			fx.Annotate(
				facades.NewDataFacade,
				fx.As(new(facades.DataService)),
			),

			// Inputs.
			fx.Annotate(
				userinteraction.NewUserInput,
				fx.As(new(userinteraction.UserInput)),
			),

			// Handlers.
			fx.Annotate(
				handler.NewBankAccountHandler,
				fx.As(new(handler.BankAccountHandler)),
			),

			fx.Annotate(
				handler.NewCategoryHandler,
				fx.As(new(handler.CategoryHandler)),
			),

			fx.Annotate(
				handler.NewOperationHandler,
				fx.As(new(handler.OperationHandler)),
			),

			fx.Annotate(
				handler.NewDefaultExportHandler,
				fx.As(new(handler.ExportHandler)),
			),

			fx.Annotate(
				handler.NewDefaultImportHandler,
				fx.As(new(handler.ImportHandler)),
			),

			// Factory.
			factories.NewBankFactory,
		),
		fx.WithLogger(func() fxevent.Logger {
			return fxevent.NopLogger
		}),
		fx.Invoke(run),
	).Run()
}

func run(
	bankAccountHandler handler.BankAccountHandler,
	categoryHandler handler.CategoryHandler,
	operationHandler handler.OperationHandler,
	analyticsFacade *facades.AnalyticsFacade,
	exportHadler handler.ExportHandler,
	importHandler handler.ImportHandler,
	repo domain.Repository,
) {
	for {
		prompt := promptui.Select{
			Label: "Select an action",
			Items: []string{
				"Work with domain model",
				"Analytics",
				"Import data",
				"Export data",
				"Manage data",
				"Exit",
			},
		}

		_, result, err := prompt.Run()
		if err != nil {
			log.Fatalf("Selection error: %v", err)
		}

		switch result {
		case "Work with domain model":
			handleDomainModel(bankAccountHandler, categoryHandler, operationHandler, repo)
		case "Analytics":
			handleAnalytics(analyticsFacade)
		case "Import data":
			handleDataImport(importHandler)
		case "Export data":
			handleDataExport(exportHadler)
		case "Manage data":
			handleDataManagement(analyticsFacade)
		case "Exit":
			os.Exit(0)
			return
		}
	}

}

func handleDomainModel(
	bankAccountHandler handler.BankAccountHandler,
	categoryHandler handler.CategoryHandler,
	operationHandler handler.OperationHandler,
	repo domain.Repository,
) {
	prompt := promptui.Select{
		Label: "Select an action for the domain model",
		Items: []string{
			"Create account",
			"Edit account",
			"Delete account",
			"Create category",
			"Edit category",
			"Delete category",
			"Create operation",
			"Edit operation",
			"Delete operation",
			"Back",
		},
	}

	_, result, err := prompt.Run()
	if err != nil {
		log.Fatalf("Selection error: %v", err)
	}

	switch result {
	case "Create account":
		if err := bankAccountHandler.CreateBankAccount(); err != nil {
			log.Fatalf("Error creating account: %v", err)
		}
	case "Edit account":
		if err := bankAccountHandler.UpdateBankAccount(); err != nil {
			log.Fatalf("Error editing account: %v", err)
		}
	case "Delete account":
		if err := bankAccountHandler.DeleteBankAccount(); err != nil {
			log.Fatalf("Error deleting account: %v", err)
		}
	case "Create category":
		if err := categoryHandler.CreateCategory(); err != nil {
			log.Fatalf("Error creating category: %v", err)
		}
	case "Edit category":
		if err := categoryHandler.UpdateCategory(); err != nil {
			log.Fatalf("Error editing category: %v", err)
		}
	case "Delete category":
		if err := categoryHandler.DeleteCategory(); err != nil {
			log.Fatalf("Error deleting category: %v", err)
		}
	case "Create operation":
		if err := operationHandler.CreateOperation(); err != nil {
			log.Fatalf("Error creating operation: %v", err)
		}
	case "Edit operation":
		if err := operationHandler.UpdateOperation(); err != nil {
			log.Fatalf("Error editing operation: %v", err)
		}
	case "Delete operation":
		if err := operationHandler.DeleteOperation(); err != nil {
			log.Fatalf("Error deleting operation: %v", err)
		}
	case "Back":
		return
	}
}

func handleAnalytics(analyticsFacade *facades.AnalyticsFacade) {
	prompt := promptui.Select{
		Label: "Select an action for analytics",
		Items: []string{
			"Calculate income and expense difference",
			"Calculate balance",
			"Back",
		},
	}

	_, result, err := prompt.Run()
	if err != nil {
		log.Fatalf("Selection error: %v", err)
	}

	switch result {
	case "Calculate income and expense difference":
		val, err := analyticsFacade.CalculateBalanceDifference()
		if err != nil {
			log.Fatalf("Error calculating difference: %v", err)
		}

		fmt.Printf("Difference between income and expenses: %v\n", val)
	case "Calculate balance":
		val, err := analyticsFacade.CalculateBalance()
		if err != nil {
			log.Fatalf("Error calculating balance: %v", err)
		}

		fmt.Printf("Balance: %v\n", val)
	case "Back":
		return
	}
}

func handleDataImport(dataImportFacade handler.ImportHandler) {
	prompt := promptui.Select{
		Label: "Select an action for data import/export",
		Items: []string{
			"Import JSON",
			"Import YAML",
			"Back",
		},
	}

	_, result, err := prompt.Run()
	if err != nil {
		log.Fatalf("Selection error: %v", err)
	}

	switch result {
	case "Import JSON":
		handleDataImportJSON(dataImportFacade)
	case "Import YAML":
		handleDataImportYAML(dataImportFacade)
	case "Back":
		return
	}
}

func handleDataImportJSON(dataImportHandler handler.ImportHandler) {
	prompt := promptui.Select{
		Label: "Select an object to import",
		Items: []string{
			"Account",
			"Category",
			"Operation",
			"Back",
		},
	}

	_, result, err := prompt.Run()
	if err != nil {
		log.Fatalf("Selection error: %v", err)
	}

	switch result {
	case "Account":
		if err := dataImportHandler.ImportJSON(handler.ImportBankAccount); err != nil {
			log.Fatalf("Error importing accounts: %v", err)
		}
	case "Category":
		if err := dataImportHandler.ImportJSON(handler.ImportCategory); err != nil {
			log.Fatalf("Error importing categories: %v", err)
		}
	case "Operation":
		if err := dataImportHandler.ImportJSON(handler.ImportOperation); err != nil {
			log.Fatalf("Error importing operations: %v", err)
		}
	case "Back":
		return
	}
}

func handleDataImportYAML(dataImportHandler handler.ImportHandler) {
	prompt := promptui.Select{
		Label: "Select an object to import",
		Items: []string{
			"Account",
			"Category",
			"Operation",
			"Back",
		},
	}

	_, result, err := prompt.Run()
	if err != nil {
		log.Fatalf("Selection error: %v", err)
	}

	switch result {
	case "Account":
		if err := dataImportHandler.ImportYAML(handler.ImportBankAccount); err != nil {
			log.Fatalf("Error importing accounts: %v", err)
		}
	case "Category":
		if err := dataImportHandler.ImportYAML(handler.ImportCategory); err != nil {
			log.Fatalf("Error importing categories: %v", err)
		}
	case "Operation":
		if err := dataImportHandler.ImportYAML(handler.ImportOperation); err != nil {
			log.Fatalf("Error importing operations: %v", err)
		}
	case "Back":
		return
	}
}

func handleDataExport(dataExportHandler handler.ExportHandler) {
	prompt := promptui.Select{
		Label: "Select an action for data import/export",
		Items: []string{
			"Export JSON",
			"Export YAML",
			"Back",
		},
	}

	_, result, err := prompt.Run()
	if err != nil {
		log.Fatalf("Selection error: %v", err)
	}

	switch result {
	case "Export JSON":
		handleDataExortJSON(dataExportHandler)
	case "Export YAML":
		handleDataExortYAML(dataExportHandler)
	case "Back":
		return
	}
}

func handleDataExortJSON(dataExportHandler handler.ExportHandler) {
	prompt := promptui.Select{
		Label: "Select an object to export",
		Items: []string{
			"Account",
			"Category",
			"Operation",
			"Back",
		},
	}

	_, result, err := prompt.Run()
	if err != nil {
		log.Fatalf("Selection error: %v", err)
	}

	switch result {
	case "Account":
		if err := dataExportHandler.ExportJSON(handler.ExportBankAccount); err != nil {
			log.Fatalf("Error exporting accounts: %v", err)
		}
	case "Category":
		if err := dataExportHandler.ExportJSON(handler.ExportCategory); err != nil {
			log.Fatalf("Error exporting categories: %v", err)
		}
	case "Operation":
		if err := dataExportHandler.ExportJSON(handler.ExportOperation); err != nil {
			log.Fatalf("Error exporting operations: %v", err)
		}
	case "Back":
		return
	}
}

func handleDataExortYAML(dataExportFacade handler.ExportHandler) {
	prompt := promptui.Select{
		Label: "Select an object to export",
		Items: []string{
			"Account",
			"Category",
			"Operation",
			"Back",
		},
	}

	_, result, err := prompt.Run()
	if err != nil {
		log.Fatalf("Selection error: %v", err)
	}

	switch result {
	case "Account":
		if err := dataExportFacade.ExportYAML(handler.ExportBankAccount); err != nil {
			log.Fatalf("Error exporting accounts: %v", err)
		}
	case "Category":
		if err := dataExportFacade.ExportYAML(handler.ExportCategory); err != nil {
			log.Fatalf("Error exporting categories: %v", err)
		}
	case "Operation":
		if err := dataExportFacade.ExportYAML(handler.ExportOperation); err != nil {
			log.Fatalf("Error exporting operations: %v", err)
		}
	case "Back":
		return
	}
}

func handleDataManagement(analyticsFacade *facades.AnalyticsFacade) {
	prompt := promptui.Select{
		Label: "Select an action for data management",
		Items: []string{
			"Recalculate balances",
			"Back",
		},
	}

	_, result, err := prompt.Run()
	if err != nil {
		log.Fatalf("Selection error: %v", err)
	}

	switch result {
	case "Recalculate balances":
		if err := analyticsFacade.RecalculateBalances(); err != nil {
			log.Fatalf("Error recalculating balances: %v", err)
		}
	case "Back":
		return
	}
}
