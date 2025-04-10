package domain

type Repository interface {
	// Save a bank account in the repository.
	SaveBankAccount(account *BankAccount) error

	// Get a bank account from the repository.
	GetBankAccount(id int) (*BankAccount, error)

	// Delete a bank account from the repository.
	DeleteBankAccount(id int) error

	// Update a bank account in the repository.
	UpdateBankAccount(account *BankAccount) error

	// Get all bank accounts from the repository.
	GetAllBankAccounts() ([]*BankAccount, error)

	// Save a category in the repository.
	SaveCategory(category *Category) error

	// Get a category from the repository.
	GetCategory(id int) (*Category, error)

	// Delete a category from the repository.
	DeleteCategory(id int) error

	// Update a category in the repository.
	UpdateCategory(category *Category) error

	// Get all categories from the repository.
	GetAllCategories() ([]*Category, error)

	// Save an operation in the repository.
	SaveOperation(op *Operation) error

	// Get an operation from the repository.
	GetOperation(id int) (*Operation, error)

	// Delete an operation from the repository.
	DeleteOperation(id int) error

	// Update an operation in the repository.
	UpdateOperation(op *Operation) error

	// Get all operations from the repository.
	GetAllOperations() ([]*Operation, error)

	// Get all operations by account ID from the repository.
	GetOperationsByAccountID(accountID int) ([]*Operation, error)
}
