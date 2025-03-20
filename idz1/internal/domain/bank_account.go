package domain

type BankAccount struct {
	ID      int     `json:"id" yaml:"id"`
	Name    string  `json:"name" yaml:"name"`
	Balance float64 `json:"balance" yaml:"balance"`
}

func (b *BankAccount) Accept(visitor ExportVisitor) error {
	return visitor.VisitBankAccount(b)
}
