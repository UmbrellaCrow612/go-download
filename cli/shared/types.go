package shared

// Options passed to the program
type Options struct {
	// The url to fetch
	Url string

	// The path to download the file or folder at
	Path string

	// A name to given the download file or folder - defaults to empty string meaning keep it's orginal name
	Name string
}
