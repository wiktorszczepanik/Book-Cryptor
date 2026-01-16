package internal

import (
	"log"
	"os"
)

func GetFile(filePath string) *os.File {
	file, error := os.Open(filePath)
	if error != nil {
		log.Fatal(error)
	}
	return file
}
