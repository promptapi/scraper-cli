![Version](https://img.shields.io/badge/version-0.1.0-orange.svg)
![Go](https://img.shields.io/badge/go-1.15.1-black.svg)
[![Documentation](https://godoc.org/github.com/promptapi/scraper-cli?status.svg)](https://pkg.go.dev/github.com/promptapi/scraper-cli)
[![Go Report Card](https://goreportcard.com/badge/github.com/promptapi/scraper-cli)](https://goreportcard.com/report/github.com/promptapi/scraper-cli)
[![Build Status](https://travis-ci.org/promptapi/scraper-cli.svg?branch=main)](https://travis-ci.org/promptapi/scraper-cli)

# Prompt API - Scraper - Command Line Interface

Do you like [Scraper API][scraper-api]? Don’t you like to access this beauty
from your terminal?

---

## Requirements

1. You need to signup for [Prompt API][promptapi-signup]
1. You need to subscribe [scraper api][scraper-api], test drive is **free!!!**
1. You need to set `PROMPTAPI_TOKEN` environment variable after subscription.

If you have `golang` environment installed, you can fetch via

```bash
$ go get -u github.com/promptapi/scraper-cli
```

Or you can download binaries from [releases][releases] page

---

## Usage

```bash
$ scraper-cli -h

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
```

---

## License

This project is licensed under MIT

---

## Contributer(s)

* [Prompt API](https://github.com/promptapi) - Creator, maintainer

---

## Contribute

All PR’s are welcome!

1. `fork` (https://github.com/promptapi/scraper-cli/fork)
1. Create your `branch` (`git checkout -b my-feature`)
1. `commit` yours (`git commit -am 'Add awesome features...'`)
1. `push` your `branch` (`git push origin my-feature`)
1. Than create a new **Pull Request**!

This project is intended to be a safe,
welcoming space for collaboration, and contributors are expected to adhere to
the [code of conduct][coc].

---

[scraper-api]:      https://promptapi.com/marketplace/description/scraper-api
[promptapi-signup]: https://promptapi.com/#signup-form
[coc]:              https://github.com/promptapi/scraper-cli/blob/main/CODE_OF_CONDUCT.md
[releases]:         https://github.com/promptapi/scraper-cli/releases