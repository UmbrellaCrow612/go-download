package shared

// Options passed to the program
type Options struct {
	// The url to fetch
	Url string

	// The path to download the file or folder at - resolvces to absolute path
	DownloadPath string

	// If it should print to stdout
	Verbose bool
}
