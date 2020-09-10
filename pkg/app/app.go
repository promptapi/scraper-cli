package app

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/url"
	"os"

	scraper "github.com/promptapi/scraper-go"
)

const (
	version = "0.1.0"
)

var (
	optVersionInformation *bool
	optURL                *string
	optAPIToken           *string
	optCountryCode        *string
	optAuthUsername       *string
	optAuthPassword       *string
	optCookie             *string
	optReferer            *string
	optSelector           *string

	usage = `
usage: scraper-cli [flags...]

  scraper-cli is a command-line interface for Prompt API's Scraper API. Details
  can be found:

      https://promptapi.com/marketplace/description/scraper-api

  you need to signup for Prompt API to get your PROMPTAPI_TOKEN. you can signup
  from:

      https://promptapi.com/#signup-form

  application looks for PROMPTAPI_TOKEN environment variable. if you pass 
  "token" flag, this will override environment variable lookup.

  required flag(s):

  -url          web url/address to scrape


  optional flags:

  -country      2 character country code
  -token        promptapi apikey instead of PROMPTAPI_TOKEN env-var
  -username     for HTTP Realm auth username
  -password     for HTTP Realm auth password
  -cookie       URL Encoded cookie header
  -referer      HTTP referer header
  -selector     CSS style selector path such as: a.btn div li
  -version      display version information
  -help, -h     display help


  examples:

  $ scraper-cli -help
  $ scraper-cli -url "https://promptapi.com"
  $ scraper-cli -url "https://promptapi.com" -country "EE"
  $ scraper-cli -url "https://promptapi.com" -country "EE" -selector "a.btn div li"

  $ PROMPTAPI_TOKEN="your-api-key" scraper-cli -url "https://promptapi.com"
  $ scraper-cli -url "https://promptapi.com" -token "your-api-key"


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
	optVersionInformation = flag.Bool("version", false, "display version information")
	optURL = flag.String("url", "", "web url/address to scrape")
	optAPIToken = flag.String("token", "n/a", "use this flag to override PROMPTAPI_TOKEN environment variable")
	optCountryCode = flag.String("country", "n/a", "2 character country code")
	optAuthUsername = flag.String("username", "n/a", "for HTTP Realm auth username")
	optAuthPassword = flag.String("password", "n/a", "for HTTP Realm auth password")
	optCookie = flag.String("cookie", "n/a", "URL Encoded cookie header")
	optReferer = flag.String("referer", "n/a", "HTTP referer header")
	optSelector = flag.String("selector", "n/a", "CSS style selector path such as: a.btn div li")

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

	if *optAPIToken != "n/a" {
		os.Setenv("PROMPTAPI_TOKEN", *optAPIToken)
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
	}
	if *optAuthUsername != "n/a" {
		params.AuthUsername = *optAuthUsername
	}
	if *optAuthPassword != "n/a" {
		params.AuthPassword = *optAuthPassword
	}
	if *optCookie != "n/a" {
		params.Cookie = *optCookie
	}
	if *optReferer != "n/a" {
		params.Referer = *optReferer
	}
	if *optSelector != "n/a" {
		params.Selector = *optSelector
	}

	result := new(scraper.Result)
	err := s.Scrape(params, result)
	if err != nil {
		return err
	}

	outData := result.Data

	if len(result.DataSelector) > 0 {
		buffer := new(bytes.Buffer)
		encoder := json.NewEncoder(buffer)
		encoder.SetEscapeHTML(false)
		err := encoder.Encode(result.DataSelector)
		if err != nil {
			return err
		}
		outData = buffer.String()
	}

	fmt.Fprintf(c.Out, outData)
	return nil
}
