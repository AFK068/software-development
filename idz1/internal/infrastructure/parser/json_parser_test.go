package parser_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/AFK068/bot/internal/domain"
	"github.com/AFK068/bot/internal/infrastructure/parser"
)

func TestJSONParser(t *testing.T) {
	jsonParser := parser.NewJSON()

	data := []byte(`[
        {"id": 1, "name": "Test Account", "balance": 1000},
        {"id": 2, "name": "Another Account", "balance": 2000}
    ]`)

	var accounts []*domain.BankAccount
	err := jsonParser.Parse(data, &accounts)
	assert.NoError(t, err)
	assert.Len(t, accounts, 2)
	assert.Equal(t, 1, accounts[0].ID)
	assert.Equal(t, "Test Account", accounts[0].Name)
	assert.Equal(t, 1000.0, accounts[0].Balance)
	assert.Equal(t, 2, accounts[1].ID)
	assert.Equal(t, "Another Account", accounts[1].Name)
	assert.Equal(t, 2000.0, accounts[1].Balance)
}
