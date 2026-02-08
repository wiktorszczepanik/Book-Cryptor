package file

import (
	"errors"
	"log"
	"bufio"
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

func GetKeyFileExt(file *os.File) (string, error) {
	ext := filepath.Ext(file.Name())
	if !(ext == "txt" || ext == "pdf" || ext == "epub") {
		return "", errors.New("Incorrect key file extension.")
	}
	return ext, nil
}

func CheckInputFile(file *os.File) error {
	ext := filepath.Ext(file.Name())
	if ext != "txt" {
		return errors.New("Incorrect input file extension.")
	}
	return nil
}

func CollectTxtRuneSet(txtFile *os.File) map[rune]bool {
	runeSet := make(map[rune]bool)
	scanner := bufio.NewScanner(txtFile)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		for _, letter := range word {
			runeSet[letter] = true
		}
	}
	return runeSet
}

func CollectPdfRuneSet(txtFile *os.File) map[rune]bool {
	return nil
}

func CollectEpubRuneSet(txtFile *os.File) map[rune]bool {
	return nil
}