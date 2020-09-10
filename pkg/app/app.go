package app

import (
	"fmt"
	"io"
	"os"
)

// CLIApplication represents app structure
type CLIApplication struct {
	Out io.Writer
}

// NewCLIApplication creates new CLIApplication instance
func NewCLIApplication() *CLIApplication {
	return &CLIApplication{
		Out: os.Stdout,
	}
}

// Run executes main application
func (c *CLIApplication) Run() error {
	fmt.Fprintf(c.Out, "Hello world")
	return nil
}
