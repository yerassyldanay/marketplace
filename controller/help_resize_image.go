package controller

import (
	"bytes"
	"errors"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

var Resize_and_store_image = func(resize_values map[string][]int, fileInBytes []byte, directory_name string,
		file_name string, file_format string, ) (error) {
	imageImage, _, err := image.Decode(bytes.NewReader(fileInBytes))

	if err != nil {
		return err
	}

	var jpeg_options jpeg.Options
	jpeg_options.Quality = 1

	for keyFormat, valueFormat := range resize_values {
		var tempFileDir string

		/*
			file_name = abc
			file_format = .png

			default: abc.png
			otherwise: abc_n_x_m.png
		 */
		if keyFormat == "default" {
			tempFileDir = file_name + file_format
		} else {
			tempFileDir = file_name + "_" + keyFormat + file_format
			imageImage = resize.Resize(uint(valueFormat[0]), uint(valueFormat[1]), imageImage, resize.Bilinear)
		}

		/*
			create a new file to store an image
		 */
		imageFile, err := os.OpenFile(filepath.Join(directory_name, filepath.Base(tempFileDir)),
			os.O_EXCL|os.O_WRONLY|os.O_CREATE, 0644)

		if err != nil {
			return err
		}

		//if keyFormat == "default" {
		//	_, _ = imageFile.Write(fileInBytes)
		//} else {
		if file_format == ".png" {
			_ = png.Encode(imageFile, imageImage)
		} else if file_format == ".jpeg" {
			_ = jpeg.Encode(imageFile, imageImage, &jpeg_options)
		} else {
			return errors.New("file format is not supported")
		}
	}

	return nil
}
