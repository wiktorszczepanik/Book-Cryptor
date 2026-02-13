package encrypt

import (
	"book-cryptor/inter/file"
	"book-cryptor/inter/oper"
	"bufio"
	"os"
	"strings"
)

type bealeEncryptCipherInfo struct {
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

func EncryptBeale(input, key *os.File) (string, error) {
	cipher := &bealeEncryptCipherInfo{}
	var err error
	if err = checkBeale(input, key, cipher); err != nil {
		return "", err
	}
	if err = oper.CollectInputSlice(input, &cipher.InputSlice); err != nil {
		return "", err
	}
	switch cipher.KeyFileExt {
	case ".txt":
		err = cipher.collectBealeReferenceFromTxt(key)
	case ".pdf":
		err = encryptBealeFromPdf(input, key, cipher)
	case ".epub":
		err = encryptBealeFromEpub(input, key, cipher)
	}
	if err != nil {
		return "", err
	}
	if cipher.OutputSlice, err = oper.GenerateCipher(cipher.InputSlice, cipher.KeyReferenceMap); err != nil {
		return "", err
	}
	var output string
	if output, err = oper.ConvertToString(&cipher.OutputSlice); err != nil {
		return "", err
	}
	return output, nil
}

func (cipher *bealeEncryptCipherInfo) collectBealeReferenceFromTxt(key *os.File) error {
	cipher.KeyReferenceMap = make(map[rune][]int)
	keyScanner := bufio.NewScanner(key)
	keyScanner.Split(bufio.ScanWords)
	var runeCounter int = 1
	for keyScanner.Scan() {
		firstRune := ([]rune(keyScanner.Text()))[0]
		if cipher.InputRuneSet[firstRune] {
			cipher.KeyReferenceMap[firstRune] = append(cipher.KeyReferenceMap[firstRune], runeCounter)
		}
		runeCounter++
	}
	if err := keyScanner.Err(); err != nil {
		return err
	}
	return nil
}

func encryptBealeFromPdf(input, key *os.File, cipher *bealeEncryptCipherInfo) error {
	return nil
}

func encryptBealeFromEpub(input, key *os.File, cipher *bealeEncryptCipherInfo) error {
	return nil
}

func checkBeale(input, key *os.File, cipher *bealeEncryptCipherInfo) error {
	var err error
	if cipher.InputFileSize, err = file.GetRuneFileSize(input); err != nil {
		return err
	}
	if cipher.KeyFileSize, err = file.GetRuneFileSize(key); err != nil {
		return err
	}
	if cipher.InputFileExt, err = file.GetInputFileExt(input); err != nil {
		return err
	}
	if cipher.KeyFileExt, err = file.GetKeyFileExt(key); err != nil {
		return err
	}
	if cipher.InputRuneSet, cipher.OutputSize, err = file.CollectAllTxtRuneSet(input); err != nil {
		return err
	}
	switch cipher.KeyFileExt {
	case ".txt":
		err = cipher.collectBealeTxtRuneSet(key)
	case ".pdf":
		cipher.KeyRuneSet, err = collectBealePdfRuneSet(key)
	case ".epub":
		cipher.KeyRuneSet, err = collectBealeEpubRuneSet(key)
	}
	if err != nil {
		return err
	}
	if err = oper.CompareRuneSets(cipher.InputRuneSet, cipher.KeyRuneSet); err != nil {
		return err
	}
	return nil
}

func (cipher *bealeEncryptCipherInfo) collectBealeTxtRuneSet(key *os.File) error {
	cipher.KeyRuneSet = make(map[rune]bool)
	scanner := bufio.NewScanner(key)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		cipher.KeyRuneSet[[]rune(word)[0]] = true
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	key.Seek(0, 0)
	return nil
}

func collectBealePdfRuneSet(pdfFile *os.File) (map[rune]bool, error) {
	return nil, nil
}

func collectBealeEpubRuneSet(epubFile *os.File) (map[rune]bool, error) {
	return nil, nil
}
