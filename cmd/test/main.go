package main

import (
	"fmt"
	"net/url"
	"os"
)

func isValidURL(urlStr string) bool {
	parsedURL, err := url.Parse(urlStr)
	if err != nil || parsedURL.Scheme == "" || parsedURL.Host == "" {
		return false
	}
	return true
}

func isLocalPath(path string) bool {
	_, err := os.Stat(path)
	// os.IsNotExist(err) returns true if the error is caused by a non-existing file or directory.
	return !os.IsNotExist(err)
}

func main() {
	testURLs := []string{
		"https://www.example.com",
		"ftp://example.com",
		"not_a_url",
		"http://example.com/path?query=value#fragment",
	}

	for _, urlStr := range testURLs {
		if isValidURL(urlStr) {
			fmt.Printf("%s is a valid URL\n", urlStr)
		} else {
			fmt.Printf("%s is not a valid URL\n", urlStr)
		}
	}

	testPaths := []string{
		"./existing_directory",
		"./existing_file.txt",
		"./non_existing_path",
		"/Users/uuxia/Desktop/work/code/go/go-service-framework/version.txt",
	}

	for _, path := range testPaths {
		if isLocalPath(path) {
			fmt.Printf("%s is a local path\n", path)
		} else {
			fmt.Printf("%s is not a local path\n", path)
		}
	}
}
