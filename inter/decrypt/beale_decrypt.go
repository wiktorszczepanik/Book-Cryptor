package decrypt

import (
	"book-cryptor/inter/decrypt/oper"
	"book-cryptor/inter/file"
	"bufio"
	"os"
)

type bealeDecryption interface {
	collectBealeTxtRuneMap(key *os.File) error
}

type bealeDecryptInfo struct {
	// Data structures
	InputSlice, SortedInputSlice []int
	KeyReferenceMap              map[int]rune

	// Output info
	OutputSlice []rune
	OutputText  string
}

// Not tested...
func Beale(input, key *os.File, separator string) (string, error) {
	var plaintext bealeDecryption = &bealeDecryptInfo{}
	var output string = ""
	if plaintext, ok := plaintext.(*bealeDecryptInfo); ok {
		if err := checkBeale(input, key); err != nil {
			return "", err
		}
		if err := oper.FileToSlice(input, &plaintext.InputSlice, separator); err != nil {
			return "", err
		}
		plaintext.SortedInputSlice = oper.SortSlice(&plaintext.InputSlice)
		plaintext.collectBealeTxtRuneMap(key)
		plaintext.OutputSlice = *oper.ReferenceMapToSlice(&plaintext.InputSlice, plaintext.KeyReferenceMap)
		plaintext.OutputText = oper.DecodedSliceToText(&plaintext.OutputSlice, &plaintext.InputSlice, plaintext.KeyReferenceMap)
		output = plaintext.OutputText
	}
	return output, nil
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

func (plaintext *bealeDecryptInfo) collectBealeTxtRuneMap(key *os.File) error {
	plaintext.KeyReferenceMap = make(map[int]rune, len(plaintext.InputSlice))
	scanner := bufio.NewScanner(key)
	scanner.Split(bufio.ScanWords)
	wordCounter, tokenSlice := 1, 0
	for scanner.Scan() && tokenSlice < len(plaintext.SortedInputSlice) {
		letter := []rune(scanner.Text())[0]
		if wordCounter == plaintext.SortedInputSlice[tokenSlice] {
			plaintext.KeyReferenceMap[plaintext.SortedInputSlice[tokenSlice]] = letter
			tokenSlice++
		}
		wordCounter++
	}
	key.Seek(0, 0)
	return nil
}
