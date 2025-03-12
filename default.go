package insomnium

import (
	"encoding/json"
	"io"
	"os"
)

// NewDefaultUnmarshaler returns an instance of the default FileReader used by
// the *Insomnium instance.
func NewDefaultFileReader() FileReader {
	return &DefaultFileReader{}
}

// NewDefaultUnmarshaler returns an instance of the default Unmarshaler used by
// the *Insomnium instance.
func NewDefaultUnmarshaler() Unmarshaler {
	return &DefaultUnmarshaler{}
}

// DefaultFileReader is an implementation of FileReader that uses os.Open under
// the hood.
type DefaultFileReader struct{}

// ReadFile implements FileReader.
func (dfr *DefaultFileReader) ReadFile(path string) (io.ReadCloser, error) {
	return os.Open(path)
}

// DefaultUnmarshaler is an implementation of Unmarshaler that uses
// json.Unmarshal under the hood.
type DefaultUnmarshaler struct{}

// Unmarshal implements Unmarshaler.
func (du *DefaultUnmarshaler) Unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}
