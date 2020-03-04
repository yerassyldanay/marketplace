package models

import (
	"os"
	"path/filepath"
)

func Remove_files_by_pattern (directory_name string, filename_pattern string, file_format string) error {
	files, err := filepath.Glob(directory_name + "/" + filename_pattern + file_format)
	if err != nil {
		return err
	}

	//fmt.Println("Found files: ", files)
	for _, f := range files {
		//fmt.Println("Removed the file: ", f)
		if err := os.Remove(f); err != nil {
			continue
		}
	}
	return nil
}
