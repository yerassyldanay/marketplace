package main

import (
	"fmt"
	"marketplace/models"
	"marketplace/utils"
	"os"
	"path/filepath"
)

func main56() {
	/*
		At the end of the main function, close the .log file
		(for more information, please, refer to utils.app_logger.go)
	*/
	defer utils.GetLogFile().Close()

	/*
		At the end of the function, close connection with database
	*/
	defer models.GetDB().Close()

	var a = "./" + utils.Folder_where_documents_are_stored + "/" + "f16E9lMJEuSILKHeSzsA4Lfwf9asDSvWDrcAOg1G2euQkxD8QD" + "*"
	fmt.Println("a: ", a)
	files, err := filepath.Glob(a)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("files: ", files)
	for _, f := range files {
		if err := os.Remove(f); err != nil {
			fmt.Println("Delete a course logo. Error: " + err.Error())
			continue
		}
	}
}