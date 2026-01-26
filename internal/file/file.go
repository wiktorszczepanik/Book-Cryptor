package file

import (
	"errors"
	"log"
	"os"
	"path/filepath"
)

func GetFile(filePath string) *os.File {
	file, error := os.Open(filePath)
	if error != nil {
		log.Fatal(error)
	}
	return file
}

func GetKeyFileExt(file *os.File) string {
	ext := filepath.Ext(file.Name())
	if !(ext == "txt" || ext == "pdf" || ext == "epub") {
		log.Fatal(errors.New("Incorrect key file extension."))
	}
	return ext
}

func GetInputFileExt(file *os.File) string {
	ext := filepath.Ext(file.Name())
	if ext != "txt" {
		log.Fatal(errors.New("Incorrect input file extension."))
	}
	return ext
}
