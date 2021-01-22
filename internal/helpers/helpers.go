package helpers

import (
	"os"
	"path"
)

// Mkdir needs a better description.
func Mkdir(perm os.FileMode, x ...string) error {
	path := path.Join(x...)

	err := os.Mkdir(path, perm)
	return err
}

// MkdirAll needs a better description.
func MkdirAll(perm os.FileMode, x ...string) error {
	path := path.Join(x...)

	err := os.MkdirAll(path, perm)
	return err
}
