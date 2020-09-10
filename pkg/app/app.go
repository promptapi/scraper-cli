package app

import (
	"flag"
	"fmt"
	"io"
	"net/url"
	"os"

	scraper "github.com/promptapi/scraper-go"
)

const (
	version = "0.0.0"
)

var (
	optVersionInformation *bool
	optURL                *string
	optAPIToken           *string
	optCountryCode        *string

	usage = `
usage: scraper-cli [options...]

  scraper-cli is a command-line interface for Prompt API's Scraper API. Details
  can be found:

  https://promptapi.com/marketplace/description/scraper-api

  you need to signup for Prompt API to get your PROMPTAPI_TOKEN. you can signup
  from:

  https://promptapi.com/#signup-form

  application looks for PROMPTAPI_TOKEN environment variable. if you pass 
  "token" flag, this will override environment variable lookup.

  example token usage:

      $ PROMPTAPI_TOKEN="your-api-key" scraper-cli -url "https://promptapi.com" # or
      $ scraper-cli -url "https://promptapi.com" -token "your-api-key"

  options:

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
	optAPIToken = flag.String("token", "n/a", "use this flag to override PROMPTAPI_TOKEN environment variable")
	optCountryCode = flag.String("country", "n/a", "2 character country code.")

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

	if err := c.Validate(); err != nil {
		return err
	}

	return c.Scrape()
}

// Version returns the current version of CLIApplication
func (c *CLIApplication) Version() error {
	fmt.Fprintf(c.Out, "%s\n", version)
	return nil
}

// Validate runs validations for flags
func (c *CLIApplication) Validate() error {
	// validate given URL
	_, err := url.ParseRequestURI(*optURL)
	if err != nil {
		return err
	}

	return nil
}

// Scrape fetches given url
func (c *CLIApplication) Scrape() error {
	s := new(scraper.PromptAPI)

	params := &scraper.Params{
		URL: *optURL,
	}

	if *optCountryCode != "n/a" {
		params.Country = *optCountryCode
		fmt.Printf("%v\n", params.Country)
	}

	result := new(scraper.Result)
	err := s.Scrape(params, result)
	if err != nil {
		return err
	}
	fmt.Fprintf(c.Out, "Content-Length: %v", result.Headers["Content-Length"])
	return nil
}
