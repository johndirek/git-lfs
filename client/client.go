package gitmediaclient

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func Put(filename string) error {
	sha := filepath.Base(filename)
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", objectUrl(sha), file)
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	fmt.Printf("Sending %s from %s: %d\n", sha, filename, res.StatusCode)
	return nil
}

func Get(filename string) (io.ReadCloser, error) {
	sha := filepath.Base(filename)
	if stat, err := os.Stat(filename); err != nil || stat == nil {
		req, err := http.NewRequest("GET", objectUrl(sha), nil)
		if err != nil {
			return nil, err
		}

		req.Header.Set("Accept", "application/vnd.git-media")
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}

		fmt.Printf("Downloading %s: %d\n", sha, res.StatusCode)
		return res.Body, nil
	}

	return os.Open(filename)
}

func objectUrl(sha string) string {
	return "http://localhost:8080/objects/" + sha
}