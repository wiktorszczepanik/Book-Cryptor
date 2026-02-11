package file

import (
	"bufio"
	"errors"
	"os"
	"path/filepath"
	"unicode/utf8"
)

func GetFile(filePath string) (*os.File, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, errors.New("Can't open provided file.")
	}
	return file, nil
}

func GetFileSize(file *os.File) (int64, error) {
	info, err := file.Stat()
	if err != nil {
		return 0, err
	}
	return info.Size(), nil
}

func GetKeyFileExt(file *os.File) (string, error) {
	ext := filepath.Ext(file.Name())
	if !(ext == "txt" || ext == "pdf" || ext == "epub") {
		return "", errors.New("Incorrect key file extension.")
	}
	return ext, nil
}

func GetInputFileExt(file *os.File) (string, error) {
	ext := filepath.Ext(file.Name())
	if ext != "txt" {
		return "", errors.New("Incorrect input file extension.")
	}
	return "txt", nil
}

func GetCipherSize(input *os.File) (int64, error) {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanWords)
	var counter int64 = 0
	for scanner.Scan() {
		wordSize := utf8.RuneCountInString(scanner.Text())
		counter += int64(wordSize)
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	input.Seek(0, 0)
	return counter, nil
}

func CollectAllTxtRuneSet(txtFile *os.File) (map[rune]bool, int, error) {
	runeSet := make(map[rune]bool)
	scanner := bufio.NewScanner(txtFile)
	scanner.Split(bufio.ScanWords)
	var sizeCouner int = 0
	for scanner.Scan() {
		word := scanner.Text()
		for _, letter := range word {
			runeSet[letter] = true
			sizeCouner++
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, 0, err
	}
	txtFile.Seek(0, 0)
	return runeSet, 0, nil
}

func CollectPdfRuneSet(txtFile *os.File) map[rune]bool {
	return nil
}

func CollectEpubRuneSet(txtFile *os.File) map[rune]bool {
	return nil
}
