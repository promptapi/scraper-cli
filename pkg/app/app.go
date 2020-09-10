package app

import (
	"flag"
	"fmt"
	"io"
	"os"
)

const (
	version = "0.0.0"
)

var (
	optVersionInformation *bool

	usage = `
Usage: scraper-cli [options...]

@wip

Options:

  -h, --help            Display help :)
`
)

// CLIApplication represents app structure
type CLIApplication struct {
	Out io.Writer
}

// NewCLIApplication creates new CLIApplication instance
func NewCLIApplication() *CLIApplication {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, usage)
	}
	optVersionInformation = flag.Bool("version", false, "Current version information")

	flag.Parse()

	return &CLIApplication{
		Out: os.Stdout,
	}
}

// Run executes main application
func (c *CLIApplication) Run() error {
	if *optVersionInformation {
		return c.Version()
	}

	fmt.Fprintf(c.Out, "Hello world")
	return nil
}

// Version returns the current version of CLIApplication
func (c *CLIApplication) Version() error {
	fmt.Fprintf(c.Out, "%s\n", version)
	return nil
}
