package controller

import (
	"io/ioutil"
	"marketplace/models"
	"marketplace/utils"
	"mime/multipart"
	"net/http"
	"time"
)

var Parse_file = func (r *http.Request) (*multipart.FileHeader, []byte, string, error) {
	var str_chan = make(chan string)

	/*
		run goroutine that generates a string for us
	*/
	go models.Generate_Random_String_Provide_Characters_To_Choose(utils.Number_of_characters_in_an_image_name, str_chan)

	/*
		Creating a buffer of size 10 * 2^20
	*/
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		return &multipart.FileHeader{}, []byte(""), "", err
	}

	/*
		parse file using the field name "uploadFile"
	*/
	var file, handler, err = r.FormFile("uploadFile")
	if err != nil {
		return &multipart.FileHeader{}, []byte(""), "", err
	}
	defer file.Close()

	/*
		Creating a file_name, which will be used as a name of the entered image
		Timeout 3 secs
	*/
	var file_name string
	select {
	case file_name = <- str_chan:
		break
	case <- time.After(time.Second * 3):
		return &multipart.FileHeader{}, []byte(""), "", err
	}

	/*
		converting file into bytes
	*/
	fileInBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return &multipart.FileHeader{}, []byte(""), "", err
	}

	return handler, fileInBytes, file_name, nil
}