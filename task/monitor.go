package task

import (
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

// Monitor allows writing a task monitor to the console
type Monitor interface {
	// Close destroys the monitor view
	Close()
	// GetHeader gets the header
	GetHeader() *Section
	// GetFooter gets the footer
	GetFooter() *Section
}

// CreateMonitor creates a Monitor
func CreateMonitor(root *Task) Monitor {
	if terminal.IsTerminal(int(os.Stdout.Fd())) {
		return createInteractiveMonitor(root)
	}
	return createNoninteractiveMonitor(root)
}
