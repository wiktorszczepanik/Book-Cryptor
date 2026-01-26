package encrypt

import (
	"book-cryptor/internal/file"
	"bufio"
	"os"
)

func EncryptBeale(input *os.File, key *os.File) string {
	inputFileExt := file.GetInputFileExt(input)
	keyFileExt := file.GetKeyFileExt(key)

	inputRuneSet := collectTxtRuneSet(input) // always .txt file
	var keyRuneSet map[rune]bool
	switch keyFileExt {
	case "txt":
		keyRuneSet := collectTxtRuneSet(key)
	case "pdf":
		keyRuneSet := collectPdfRuneSet(key)
	case "epub":
		keyRuneSet := collectEpubRuneSet(key)
	}
	keyRuneSet := collectKeyRuneSet(key, keyFileExt)
	keyRuneSet := make(map[rune]bool)
	// switch
}

func collectTxtRuneSet(txtFile *os.File) map[rune]bool {
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

func collectPdfRuneSet(txtFile *os.File) map[rune]bool {
	return nil
}

func collectEpubRuneSet(txtFile *os.File) map[rune]bool {
	return nil
}