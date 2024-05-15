package main

import (
	"io/fs"
	"os"
	"strings"
)

func parseStaticHtml(fileSystem fs.FS) error {
	err := fs.WalkDir(fileSystem, "static_html", func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {

			if err != nil {
				return err
			}

			data, e := os.ReadFile(path)
			if e != nil {
				return e
			}

			parts := strings.Split(path, "/")
			name := strings.Split(parts[len(parts)-1], ".")[0]
			staticHTML[name] = string(data)
		}
		return nil
	})
	return err
}