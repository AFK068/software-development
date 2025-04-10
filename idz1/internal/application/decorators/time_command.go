package decorators

import (
	"log"
	"time"

	"github.com/AFK068/bot/internal/application/commands"
)

type TimeCommand struct {
	Command commands.Command
}

func NewTimeCommand(command commands.Command) *TimeCommand {
	return &TimeCommand{
		Command: command,
	}
}

func (t *TimeCommand) Execute() error {
	start := time.Now()
	err := t.Command.Execute()
	elapsed := time.Since(start)

	log.Printf("Execution time: %v", elapsed)

	return err
}
