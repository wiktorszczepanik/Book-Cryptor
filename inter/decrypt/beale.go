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
	// run basic checks on params
	if err := oper.EncryptedFileToSlice(input, &plaintext.InputSlice, separator); err != nil {
		return "", err
	}
	// sort input slice
	// loop over sorted slice and create map: map[slice int] = 'character'/'rune'
	// loop over unsorted slice and decode message

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
