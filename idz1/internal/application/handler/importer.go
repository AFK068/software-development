package handler

import (
	"log"

	"github.com/AFK068/bot/internal/application/facades"
	"github.com/AFK068/bot/internal/infrastructure/importer"
	"github.com/AFK068/bot/internal/infrastructure/parser"
	"github.com/AFK068/bot/internal/infrastructure/userinteraction"
)

type ImportObjectType string

const (
	ImportCategory    ImportObjectType = "category"
	ImportBankAccount ImportObjectType = "bank_account"
	ImportOperation   ImportObjectType = "operation"
)

type ImportHandler interface {
	ImportJSON(importType ImportObjectType) error
	ImportYAML(importType ImportObjectType) error
}

type DefaultImportHandler struct {
	dataFacade facades.DataService
	userInput  userinteraction.UserInput
}

func NewDefaultImportHandler(facade facades.DataService, userInput userinteraction.UserInput) *DefaultImportHandler {
	return &DefaultImportHandler{
		dataFacade: facade,
		userInput:  userInput,
	}
}

func (h *DefaultImportHandler) ImportJSON(importType ImportObjectType) error {
	jsonParser := parser.NewJSON()

	fileName, err := h.userInput.GetStringValue("Enter the file name for export (the .json suffix will be added automatically)")
	if err != nil {
		log.Fatalf("Input error: %v", err)
	}

	fileName += ".json"

	switch importType {
	case ImportCategory:
		importer := importer.NewCategoryImporter(jsonParser)
		err = h.dataFacade.ImportCategories(importer, fileName)
	case ImportBankAccount:
		importer := importer.NewBankAccountImporter(jsonParser)
		err = h.dataFacade.ImportBankAccounts(importer, fileName)
	case ImportOperation:
		importer := importer.NewOperationImporter(jsonParser)
		err = h.dataFacade.ImportOperations(importer, fileName)
	}

	return err
}

func (h *DefaultImportHandler) ImportYAML(importType ImportObjectType) error {
	yamlParser := parser.NewYAML()

	fileName, err := h.userInput.GetStringValue("Enter the file name for export (the .yaml suffix will be added automatically)")
	if err != nil {
		log.Fatalf("Input error: %v", err)
	}

	fileName += ".yaml"

	switch importType {
	case ImportCategory:
		importer := importer.NewCategoryImporter(yamlParser)
		err = h.dataFacade.ImportCategories(importer, fileName)
	case ImportBankAccount:
		importer := importer.NewBankAccountImporter(yamlParser)
		err = h.dataFacade.ImportBankAccounts(importer, fileName)
	case ImportOperation:
		importer := importer.NewOperationImporter(yamlParser)
		err = h.dataFacade.ImportOperations(importer, fileName)
	}

	return err
}
