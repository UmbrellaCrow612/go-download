package console

import (
	"fmt"
	"os"
	"time"

	"github.com/UmbrellaCrow612/go-download/cli/shared"
)

// ANSI color codes
const (
	ColorRed    = "\033[31m"
	ColorYellow = "\033[33m"
	ColorWhite  = "\033[37m"
	ColorReset  = "\033[0m"
)

// WriteLn prints a normal message to stdout with timestamp
func WriteLn(message string) {
	if !shared.Verbose {
		return
	}
	fmt.Printf("%s%s %s%s\n", ColorWhite, time.Now().Format("15:04:05"), message, ColorReset)
}

// WriteError prints an error message to stderr with timestamp in red
func WriteError(message string) {
	if !shared.Verbose {
		return
	}
	fmt.Fprintf(os.Stderr, "%s%s ERROR: %s%s\n", ColorRed, time.Now().Format("15:04:05"), message, ColorReset)
}

// WriteWarning prints a warning message to stdout with timestamp in yellow
func WriteWarning(message string) {
	if !shared.Verbose {
		return
	}
	fmt.Printf("%s%s WARNING: %s%s\n", ColorYellow, time.Now().Format("15:04:05"), message, ColorReset)
}

// ExitError prints an error message to stderr and exits the program
func ExitError(message string) {
	fmt.Fprintf(os.Stderr, "%s%s FATAL: %s%s\n", ColorRed, time.Now().Format("15:04:05"), message, ColorReset)
	os.Exit(1)
}
