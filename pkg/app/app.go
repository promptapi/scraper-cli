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
	optURL                *string

	usage = `
Usage: scraper-cli [options...]

@wip

Options:

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
		flag.PrintDefaults()
	}
	optVersionInformation = flag.Bool("version", false, "display version information")
	optURL = flag.String("url", "", "web url/address to scrape")

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

	fmt.Fprintf(c.Out, "url ?? %s\n", *optURL)
	fmt.Fprintf(c.Out, "Hello world\n\n")
	return nil
}

// Version returns the current version of CLIApplication
func (c *CLIApplication) Version() error {
	fmt.Fprintf(c.Out, "%s\n", version)
	return nil
}
