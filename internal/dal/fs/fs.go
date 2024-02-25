package fs

import (
	"io/fs"
	"os"

	"github.com/lstig/liber/internal/dal"
)

type FS struct {
	// Root directory to store files in
	Root string
	// File permission on the root directory
	Perm fs.FileMode
}

func (f *FS) Build() (dal.Operator, error) {
	if err := os.MkdirAll(f.Root, f.Perm); err != nil {
		return f, err
	}
	return f, nil
}

func (f *FS) Read(path string, p []byte) (n int, err error) {
	return 0, nil
}

func (f *FS) Write(path string, p []byte) (n int, err error) {
	return 0, nil
}

func (f *FS) Stat(path string) (meta *dal.Metadata, err error) {
	return nil, nil
}
