package decrypt

import (
	"book-cryptor/inter/oper"
	"os"
	"strings"
)

type bealeDecryptCipherInfo struct {
	// Files info
	InputFileExt, KeyFileExt string

	// Data structures
	InputSlice               []int
	InputRuneSet, KeyRuneSet map[rune]bool
	KeyReferenceMap          map[rune][]int

	// Output info
	OutputSize  int
	OutputSlice []int
	OutputText  strings.Builder
}

func DecryptBeale(input, key *os.File, separator string) (string, error) {
	plaintext := &bealeDecryptCipherInfo{}
	if err := oper.EncryptedFileToSlice(input, &plaintext.InputSlice, separator); err != nil {
		return "", err
	}

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
