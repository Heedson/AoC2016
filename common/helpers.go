package common

import (
	"bytes"
	"io"
)

// LoadStringFromReader accepts any reader and returns the string that is contained within it.
func LoadStringFromReader(reader io.Reader) (string, error) {
	buf := bytes.Buffer{}
	_, err := buf.ReadFrom(reader)
	if err != nil {
		return "", err
	}

	return string(buf.Bytes()), nil
}
