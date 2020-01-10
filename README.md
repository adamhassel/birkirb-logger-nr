# loggers-mapper-newrelic
Golang [Loggers](https://gopkg.in/birkirb/loggers.v1) mapper for [New Relic](https://github.com/newrelic/go-agent).

[![GoDoc](https://godoc.org/github.com/adamhassel/loggers-mapper-newrelic?status.svg)](https://godoc.org/github.com/adamhassel/loggers-mapper-newrelic)
[![Build Status](https://travis-ci.org/adamhassel/loggers-mapper-newrelic.svg?branch=master)](http://travis-ci.org/adamhassel/loggers-mapper-newrelic)

## Pre-recquisite

See https://gopkg.in/birkirb/loggers.v1

## Installation

    go get github.com/adamhassel/loggers-mapper-newrelic

## Usage

Assuming you are using loggers in your code, and you want to use loggers for New Relic. Start by configuring your loggers, and then pass it to the mapper and assign it as NewRelic's loggers interface.

### Example

```Go
package main

import (
	"os"

    "gopkg.in/birkirb/loggers.v1"
	newrelic "github.com/newrelic/go-agent"
    nrlog "github.com/adamhassel/loggers-mapper-newrelic"
)

// Log is my default logger.
var Log loggers.Contextual

func main() {
	var debug bool
	if testing {
		debug = true
	}
	config := newrelic.NewConfig(revel.AppName, license)
	logger := nrlog.NewLogger(Log, debug)
	config.Logger = nrlog.NewLogger(Log, debug)
	app, err = newrelic.NewApplication(config)
	if err != nil {
		panic(err)
	}
	// ... do new relic transactions
}
```
