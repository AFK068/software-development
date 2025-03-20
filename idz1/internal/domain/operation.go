package domain

import "time"

const (
	OperationTypeIncome  = "income"
	OperationTypeExpense = "expense"
)

type Operation struct {
	ID            int       `json:"id" yaml:"id"`
	Type          string    `json:"type" yaml:"type"`
	BankAccountID int       `json:"bank_account_id" yaml:"bank_account_id"`
	Amount        float64   `json:"amount" yaml:"amount"`
	Date          time.Time `json:"date" yaml:"date"`
	Description   string    `json:"description" yaml:"description"`
	CategoryID    int       `json:"category_id" yaml:"category_id"`
}

func (o *Operation) Accept(visitor ExportVisitor) error {
	return visitor.VisitOperation(o)
}
