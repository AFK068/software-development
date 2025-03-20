package handler

import (
	"fmt"
	"log"

	"github.com/AFK068/bot/internal/application/facades"
	"github.com/AFK068/bot/internal/infrastructure/exporter"
	"github.com/AFK068/bot/internal/infrastructure/userinteraction"
)

type ExportObjectType string

const (
	ExportCategory    ExportObjectType = "category"
	ExportBankAccount ExportObjectType = "bank_account"
	ExportOperation   ExportObjectType = "operation"
)

type ExportHandler interface {
	ExportJSON(exportType ExportObjectType) error
	ExportYAML(exportType ExportObjectType) error
}

type DefaultExportHandler struct {
	dataFacade facades.DataService
	userInput  userinteraction.UserInput
}

func NewDefaultExportHandler(facade facades.DataService, userInput userinteraction.UserInput) *DefaultExportHandler {
	return &DefaultExportHandler{
		dataFacade: facade,
		userInput:  userInput,
	}
}

func (h *DefaultExportHandler) ExportJSON(exportType ExportObjectType) error {
	jsonExport := exporter.NewJSON()

	fileName, err := h.userInput.GetStringValue("Enter the file name for export (postfix .json will be added automatically)")
	if err != nil {
		log.Fatalf("Input error: %v", err)
	}

	fileName += ".json"

	switch exportType {
	case ExportCategory:
		err = h.dataFacade.ExportCategories(jsonExport, fileName)
	case ExportBankAccount:
		err = h.dataFacade.ExportBankAccounts(jsonExport, fileName)
	case ExportOperation:
		err = h.dataFacade.ExportOperations(jsonExport, fileName)
	}

	if err != nil {
		return fmt.Errorf("failed to export data: %w", err)
	}

	return nil
}

func (h *DefaultExportHandler) ExportYAML(exportType ExportObjectType) error {
	yamlExport := exporter.NewYAML()

	fileName, err := h.userInput.GetStringValue("Enter the file name for export (postfix .yaml will be added automatically)")
	if err != nil {
		log.Fatalf("Input error: %v", err)
	}

	fileName += ".yaml"

	switch exportType {
	case ExportCategory:
		err = h.dataFacade.ExportCategories(yamlExport, fileName)
	case ExportBankAccount:
		err = h.dataFacade.ExportBankAccounts(yamlExport, fileName)
	case ExportOperation:
		err = h.dataFacade.ExportOperations(yamlExport, fileName)
	}

	if err != nil {
		return fmt.Errorf("failed to export data: %w", err)
	}

	return nil
}
