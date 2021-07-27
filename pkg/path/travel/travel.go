package travel

import (
	"io/fs"
	"path/filepath"
)

// ListPath list all directory under root, except root
func ListPath(root string) (dir []string, err error) {
	dir = make([]string, 0)
	err = filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path != root {
			dir = append(dir, path)
			return nil
		}
		return nil
	})
	return
}