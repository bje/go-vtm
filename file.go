package stingray

import (
	"fmt"
)

// fileResource represents a file resource.
type fileResource struct {
	resource
	Content []byte
	Note    string
}

func (f *fileResource) String() string {
	return fmt.Sprintf("#=-%v\n", f.Note) + string(f.Content)
}

func (f *fileResource) decode(data []byte) error {
	f.Content = data
	return nil
}

func (f *fileResource) contentType() string {
	return "application/octet-stream"
}
