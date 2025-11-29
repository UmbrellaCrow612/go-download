package main

import (
	"github.com/UmbrellaCrow612/go-download/cli/args"
	"github.com/UmbrellaCrow612/go-download/cli/console"
	"github.com/UmbrellaCrow612/go-download/cli/fetch"
	"github.com/UmbrellaCrow612/go-download/cli/shared"
)

func main() {
	options := args.Parse()

	// set global const
	shared.Verbose = options.Verbose

	err := fetch.Get(options)
	if err != nil {
		console.ExitError(err.Error())
	}
}
