package controller

import "os"

/*
	this function is created to clear files on directory
	in case something goes wrong
 */
var Remove_file_from_directory = func(directory_name string, file_name_with_format string) (error) {
	return os.Remove(directory_name + "/" + file_name_with_format)
}
