package common

import (
	"bytes"
	"io"
	"net/http"
)

// GetRawFromURL gets a file from a url and its raw content
func GetRawFromURL(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	buf := bytes.NewBuffer(make([]byte, 0, resp.ContentLength))
	_, err = io.Copy(buf, resp.Body)

	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
