package args

import (
	"net/url"
	"os"
	"path/filepath"

	"github.com/UmbrellaCrow612/go-download/cli/shared"
	"github.com/UmbrellaCrow612/go-download/cli/utils"
)

// Parses args array
func Parse() *shared.Options {
	options := &shared.Options{}

	if len(os.Args) < 3 {
		utils.ExitWithError("Error: expected at least 2 arguments: [url] [path]")
	}

	rawURL := os.Args[1]
	rawPath := os.Args[2]

	parsedURL, err := url.ParseRequestURI(rawURL)
	if err != nil {
		utils.ExitWithError("Error: invalid URL format: " + rawURL)
	}

	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		utils.ExitWithError("Error: URL must start with http or https")
	}

	options.Url = parsedURL.String()

	absPath, err := filepath.Abs(rawPath)
	if err != nil {
		utils.ExitWithError("Error: invalid path: " + rawPath)
	}

	options.Path = absPath

	for i := 3; i < len(os.Args); i++ {
		arg := os.Args[i]

		switch arg {
		case "--name":
			if i+1 >= len(os.Args) {
				utils.ExitWithError("Error: missing value for --name")
			}
			options.Name = os.Args[i+1]
			i++

		default:
			if len(arg) > 2 && arg[:2] == "--" {
				utils.ExitWithError("Error: unknown flag " + arg)
			}
			utils.ExitWithError("Error: unexpected argument: " + arg)
		}
	}

	return options
}
