package encrypt

import (
	"book-cryptor/internal/file"
	"book-cryptor/internal/operations"
	"bufio"
	"os"
)

func EncryptBeale(input, key *os.File) (string, error) {
	if err := checkBeale(input, key); err != nil {
		return "", err
	}
}

// Only for txt files (for now)
func checkBeale(input, key *os.File) error {
	if err := file.CheckInputFile(input); err != nil {
		return err
	}
	keyFileExt, err := file.GetKeyFileExt(key)
	if err != nil {
		return err
	}
	inputRuneSet, err := file.CollectTxtRuneSet(input) // always .txt file
	if err != nil {
		return err
	}
	var keyRuneSet map[rune]bool
	switch keyFileExt {
	case "txt":
		keyRuneSet, err = collectBealeTxtRuneSet(key)
	case "pdf":
		keyRuneSet, err = collectBealePdfRuneSet(key)
	case "epub":
		keyRuneSet, err = collectBealeEpubRuneSet(key)
	}
	if err != nil {
		return err
	}
	if err = operations.CompareRuneSets(inputRuneSet, keyRuneSet); err != nil {
		return err
	}
	return nil
}

// Collect first letter of every word in txt file
func collectBealeTxtRuneSet(txtFile *os.File) (map[rune]bool, error) {
	runeSet := make(map[rune]bool)
	scanner := bufio.NewScanner(txtFile)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		runeSet[[]rune(word)[0]] = true
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return runeSet, nil
}

func collectBealePdfRuneSet(txtFile *os.File) (map[rune]bool, error) {
	return nil, nil
}

func collectBealeEpubRuneSet(txtFile *os.File) (map[rune]bool, error) {
	return nil, nil
}
