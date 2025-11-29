package fetch

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/UmbrellaCrow612/go-download/cli/console"
	"github.com/UmbrellaCrow612/go-download/cli/shared"
)

// Gets the url with options
func Get(options *shared.Options) error {
	url := options.Url
	console.WriteLn("Fetching " + url)

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to fetch URL: %s", err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("server returned status %d", resp.StatusCode)
	}

	filename := extractFilename(url)
	if filename == "" {
		filename = "downloaded_file"
	}

	outputPath := filepath.Join(options.DownloadPath, filename)

	console.WriteLn("Downloading to " + outputPath)

	if err := os.MkdirAll(options.DownloadPath, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create output directory: %s", err.Error())
	}

	outFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("unable to create file: %s", err.Error())
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		return fmt.Errorf("error writing file: %s", err.Error())
	}

	return nil
}

func extractFilename(rawURL string) string {
	slash := strings.LastIndex(rawURL, "/")
	if slash == -1 || slash == len(rawURL)-1 {
		return ""
	}
	return rawURL[slash+1:]
}
