package decrypt

import (
	"book-cryptor/inter/file"
	"book-cryptor/inter/oper"
	"bufio"
	"os"
	"strings"
)

type bealeDecryptCipherInfo struct {
	// Files info
	InputFileExt, KeyFileExt   string
	InputFileSize, KeyFileSize int64

	// Data structures
	InputSlice               []rune
	InputRuneSet, KeyRuneSet map[rune]bool
	KeyReferenceMap          map[rune][]int

	// Output info
	OutputSize  int
	OutputSlice []int
	OutputText  strings.Builder
}

func decryptBeale(input, key *os.File, separator string) (string, error) {
	return "", nil
}

func (cipher *bealeDecryptCipherInfo) collectBealeReferenceFromTxt(key *os.File) error {
	return nil
}

func decryptBealeFromPdf(input, key *os.File, cipher *bealeDecryptCipherInfo) error {
	return nil
}

func decryptBealeFromEpub(input, key *os.File, cipher *bealeDecryptCipherInfo) error {
	return nil
}

func checkBeale(input, key *os.File, cipher *bealeDecryptCipherInfo) error {
	return nil
}

func (cipher *bealeDecryptCipherInfo) collectBealeTxtRuneSet(key *os.File) error {
	return nil
}

func collectBealePdfRuneSet(pdfFile *os.File) (map[rune]bool, error) {
	return nil, nil
}

func collectBealeEpubRuneSet(epubFile *os.File) (map[rune]bool, error) {
	return nil, nil
}
