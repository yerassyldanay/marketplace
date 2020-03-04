package controller

import (
	"os"
	"path/filepath"
)

var Store_document = func(directory string, filename string, file_format string, file_in_bytes []byte) (error) {
	/*
		creating a file within the indicated directory
	*/
	var tempFileDir = filename + file_format
	newDocument, err := os.OpenFile(filepath.Join(directory, filepath.Base(tempFileDir)),
		os.O_EXCL|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	/*
		writing bytes into the file (which has been created)
	*/
	_, err = newDocument.Write(file_in_bytes)
	if err != nil {
		return err
	}

	return nil
}
