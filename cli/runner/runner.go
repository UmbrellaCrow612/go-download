package runner

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/UmbrellaCrow612/go-download/cli/shared"
	"github.com/UmbrellaCrow612/go-download/cli/utils"
)

// Runs the main loop
func Run(options *shared.Options) {
	url := options.Url

	resp, err := http.Get(url)
	if err != nil {
		utils.ExitWithError("Error: failed to fetch URL: " + err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		utils.ExitWithError(fmt.Sprintf("Error: server returned status %d", resp.StatusCode))
	}

	filename := options.Name
	if filename == "" {
		filename = extractFilename(url)
		if filename == "" {
			filename = "downloaded_file"
		}
	}

	outputPath := filepath.Join(options.Path, filename)

	outFile, err := os.Create(outputPath)
	if err != nil {
		utils.ExitWithError("Error: unable to create file: " + err.Error())
	}
	defer outFile.Close()

	contentLength := resp.ContentLength
	var totalDownloaded int64 = 0
	buffer := make([]byte, 32*1024) // 32 KB chunks

	utils.PrintToStdout("Downloading to: " + outputPath)

	for {
		n, readErr := resp.Body.Read(buffer)
		if n > 0 {
			_, writeErr := outFile.Write(buffer[:n])
			if writeErr != nil {
				utils.ExitWithError("Error: failed writing to file: " + writeErr.Error())
			}
			totalDownloaded += int64(n)

			printProgress(totalDownloaded, contentLength)
		}

		if readErr == io.EOF {
			break
		}

		if readErr != nil {
			utils.ExitWithError("Error: failed reading from response body: " + readErr.Error())
		}
	}

	fmt.Println() 

	utils.PrintToStdout("Download completed successfully!")
	utils.PrintToStdout("Saved as: " + outputPath)
}

func extractFilename(rawURL string) string {
	slash := strings.LastIndex(rawURL, "/")
	if slash == -1 || slash == len(rawURL)-1 {
		return ""
	}
	return rawURL[slash+1:]
}

func printProgress(downloaded, total int64) {
	if total <= 0 {
		// Unknown size → just print bytes downloaded
		fmt.Printf("\rDownloaded %d bytes", downloaded)
		return
	}

	percent := float64(downloaded) / float64(total)
	barWidth := 40
	filled := int(percent * float64(barWidth))

	bar := "[" + strings.Repeat("=", filled) + strings.Repeat(" ", barWidth-filled) + "]"

	fmt.Printf("\r%s %.2f%% (%d/%d bytes)", bar, percent*100, downloaded, total)
}
