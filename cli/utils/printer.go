package utils

import (
	"fmt"
	"os"
)

// Print to standard output
func PrintToStdout(msg string) {
	fmt.Fprintln(os.Stdout, msg)
}

// Print to standard error
func PrintToStderr(msg string) {
	fmt.Fprintln(os.Stderr, msg)
}

// Print an error and exit the program
func ExitWithError(msg string) {
	PrintToStderr(msg)
	os.Exit(1)
}
