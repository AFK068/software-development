package factories

import (
	"time"

	"github.com/AFK068/bot/internal/domain"
)

type Factory interface {
	NewBankAccount(id int, name string, balance float64) (*domain.BankAccount, error)
	NewCategory(id int, name string, categoryType string) (*domain.Category, error)
	NewOperation(id int, typeOperation string,
		bankAccountID int, amount float64, date time.Time, description string, categoryID int) (*domain.Operation, error)
}
