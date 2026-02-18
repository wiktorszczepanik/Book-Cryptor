package decrypt

import (
	"book-cryptor/inter/decrypt/oper"
	"book-cryptor/inter/file"
	"bufio"
	"os"
)

type bealeDecryptCipherInfo struct {
	// Data structures
	InputSlice, SortedInputSlice []int
	KeyReferenceMap              map[int]rune

	// Output info
	OutputSlice []rune
	OutputText  string
}

// Not tested...
func Beale(input, key *os.File, separator string) (string, error) {
	plaintext := &bealeDecryptCipherInfo{}
	if err := checkBeale(input, key); err != nil {
		return "", err
	}
	if err := oper.FileToSlice(input, &plaintext.InputSlice, separator); err != nil {
		return "", err
	}
	plaintext.SortedInputSlice = oper.SortSlice(&plaintext.InputSlice)
	plaintext.collectBealeTxtRuneMap(key)
	plaintext.OutputSlice = *oper.ReferenceMapToSlice(&plaintext.InputSlice, plaintext.KeyReferenceMap)
	plaintext.OutputText = oper.DecodedSliceToText(&plaintext.OutputSlice)
	return plaintext.OutputText, nil
}

func checkBeale(input, key *os.File) error {
	var err error
	if err = file.CheckInputFileExt(input); err != nil {
		return err
	}
	if err = file.CheckKeyFileExt(key); err != nil {
		return err
	}
	return nil
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
	key.Seek(0, 0)
	return nil
}
