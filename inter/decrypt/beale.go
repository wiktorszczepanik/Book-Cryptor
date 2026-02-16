package decrypt

import (
	"book-cryptor/inter/file"
	"book-cryptor/inter/oper"
	"bufio"
	"os"
)

type bealeDecryptCipherInfo struct {
	// Files info
	InputFileExt, KeyFileExt string

	// Data structures
	InputSlice, SortedInputSlice []int
	KeyReferenceMap              map[int]rune

	// Output info
	OutputSlice []rune
	OutputText  string
}

// Not tested...
func DecryptBeale(input, key *os.File, separator string) (string, error) {
	plaintext := &bealeDecryptCipherInfo{}
	if err := checkBeale(input, key); err != nil {
		return "", err
	}
	if err := oper.EncryptedFileToSlice(input, &plaintext.InputSlice, separator); err != nil {
		return "", err
	}
	plaintext.SortedInputSlice = oper.GetSortedEncryptedInputSlice(&plaintext.InputSlice)
	plaintext.collectBealeTxtRuneMap(key)
	plaintext.OutputSlice = *oper.ConvertReferenceMapToSlice(&plaintext.InputSlice, plaintext.KeyReferenceMap)
	plaintext.OutputText = oper.ConvertDecodedSliceToText(&plaintext.OutputSlice)
	return plaintext.OutputText, nil
}

func (plaintext *bealeDecryptCipherInfo) collectBealeTxtRuneMap(key *os.File) error {
	plaintext.KeyReferenceMap = make(map[int]rune, len(plaintext.InputSlice))
	scanner := bufio.NewScanner(key)
	scanner.Split(bufio.ScanRunes)
	wordCounter, tokenSlice := 0, 0
	for scanner.Scan() {
		letter := []rune(scanner.Text())[0]
		if wordCounter == plaintext.SortedInputSlice[tokenSlice] {
			plaintext.KeyReferenceMap[tokenSlice] = letter
			tokenSlice++
		}
		wordCounter++
	}
	return nil
}

func checkBeale(input, key *os.File) error {
	var err error
	if err = file.CheckInputFileExt(input); err != nil {
		return err
	}
	if err = file.CheckKeyFileExt(key); err != nil {
		return nil
	}
	return nil
}
