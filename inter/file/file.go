package file

import (
	"errors"
	"os"
	"path/filepath"
)

func GetEssensialFiles(inputPath, keyPath string) (*os.File, *os.File, error) {
	var inputFile, keyFile *os.File
	var err error
	if inputFile, err = getFile(inputPath); err != nil {
		return nil, nil, err
	}
	if keyFile, err = getFile(keyPath); err != nil {
		return nil, nil, err
	}
	return inputFile, keyFile, nil
}

func getFile(filePath string) (*os.File, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, errors.New("Can't open provided file.")
	}
	return file, nil
}

func GetKeyFileExt(file *os.File) (string, error) {
	ext := filepath.Ext(file.Name())
	if !(ext == ".txt" || ext == ".pdf" || ext == ".epub") {
		return "", errors.New("Incorrect key file extension.")
	}
	return ext, nil
}

func CheckKeyFileExt(file *os.File) error {
	_, err := GetKeyFileExt(file)
	if err != nil {
		return err
	}
	return nil
}

func CheckInputFileExt(file *os.File) error {
	ext := filepath.Ext(file.Name())
	if ext != ".txt" {
		return errors.New("Incorrect input file extension.")
	}
	return nil
}

func SaveOutput(filePath string, cipherText string) error {
	cipherText += "\n"
	var file *os.File
	var err error
	if file, err = os.Create(filePath); err != nil {
		return err
	}
	defer file.Close()
	if _, err = file.WriteString(cipherText); err != nil {
		return err
	}
	return nil
}
