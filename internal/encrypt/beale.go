package encrypt

import (
	"book-cryptor/internal/file"
	"errors"
	"maps"
	"os"
)

func EncryptBeale(input *os.File, key *os.File) (string, error) {
	inputFileInfo:= file.CheckInputFile(input)
	if inputFileInfo != nil {
		return "", inputFileInfo
	}
	keyFileExt, err := file.GetKeyFileExt(key)
	if err != nil {
		return "", err
	}

	inputRuneSet := file.CollectTxtRuneSet(input) // always .txt file
	var keyRuneSet map[rune]bool
	switch keyFileExt {
	case "txt":
		keyRuneSet := file.CollectTxtRuneSet(key)
	case "pdf":
		keyRuneSet := file.CollectPdfRuneSet(key)
	case "epub":
		keyRuneSet := file.CollectEpubRuneSet(key)
	}
	// if maps.Equal(inputRuneSet, keyRuneSet) {
	// 	return "", errors.New("")
	// }

	// keyRuneSet := collectKeyRuneSet(key, keyFileExt)
	// keyRuneSet := make(map[rune]bool)
	// switch
}
