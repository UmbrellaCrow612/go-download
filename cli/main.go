package main

import (
	"github.com/UmbrellaCrow612/go-download/cli/args"
	"github.com/UmbrellaCrow612/go-download/cli/runner"
)

func main() {
	options := args.Parse()
	runner.Run(options)
}
