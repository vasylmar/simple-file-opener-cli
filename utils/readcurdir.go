package utils

import "os"

type CurrendDirItems struct {
	items []os.DirEntry
}

func (c *CurrendDirItems) ReadCurDir(path string) error {
	output, err := os.ReadDir(path)
	if err != nil {
		return err
	}
	c.items = output
	return nil
}

func (c *CurrendDirItems) GetFilesFromCurDir() []string {
	var files []string
	for _, item := range c.items {
		if !item.IsDir() {
			files = append(files, item.Name())
		}
	}
	return files
}
