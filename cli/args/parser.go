package args

import (
	"net/url"
	"os"
	"path/filepath"

	"github.com/UmbrellaCrow612/go-download/cli/console"
	"github.com/UmbrellaCrow612/go-download/cli/shared"
)

// Parses args array
func Parse() *shared.Options {
	options := &shared.Options{
		Verbose: false,
	}

	if len(os.Args) < 3 {
		console.ExitError("expected at least 2 arguments: [url] [downloadPath] [..flags..]")
	}

	rawURL := os.Args[1]
	rawPath := os.Args[2]

	parsedURL, err := url.ParseRequestURI(rawURL)
	if err != nil {
		console.ExitError("invalid URL format: " + rawURL)
	}

	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		console.ExitError("URL must start with http or https")
	}

	options.Url = parsedURL.String()

	absPath, err := filepath.Abs(rawPath)
	if err != nil {
		console.ExitError("invalid path: " + rawPath)
	}
	options.DownloadPath = absPath

	for _, arg := range os.Args[3:] {
		switch arg {
		case "--verbose", "-v":
			options.Verbose = true
		default:
			if len(arg) > 0 && arg[0] == '-' {
				console.ExitError("unknown flag: " + arg)
			}
		}
	}

	return options
}
