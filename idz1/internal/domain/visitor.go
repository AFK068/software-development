package domain

type ExportVisitor interface {
	VisitBankAccount(account *BankAccount) error
	VisitCategory(category *Category) error
	VisitOperation(operation *Operation) error
}
