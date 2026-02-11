package encrypt

import (
	"book-cryptor/inter/file"
	"book-cryptor/inter/operations"
	"bufio"
	"crypto/rand"
	"fmt"
	"math/big"
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
	if err = collectInputSlice(input, cipher); err != nil {
		return "", err
	}
	switch cipher.KeyFileExt {
	case ".txt":
		err = collectBealeReferenceFromTxt(key, cipher)
	case ".pdf":
		err = encryptBealeFromPdf(input, key, cipher)
	case ".epub":
		err = encryptBealeFromEpub(input, key, cipher)
	}
	if err != nil {
		return "", err
	}
	if err = generateBealeCipher(cipher); err != nil {
		return "", err
	}
	var output string
	if output, err = convertToString(cipher); err != nil {
		return "", err
	}
	return output, nil
}

func collectInputSlice(input *os.File, cipher *bealeEncryptCipherInfo) error {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		for _, r := range word {
			cipher.InputSlice = append(cipher.InputSlice, r)
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func generateBealeCipher(cipher *bealeEncryptCipherInfo) error {
	for _, v := range cipher.InputSlice {
		selection := cipher.KeyReferenceMap[v]
		selectionLength := int64(len(selection))
		random, err := rand.Int(rand.Reader, big.NewInt(selectionLength))
		if err != nil {
			return err
		}
		cipher.OutputSlice = append(cipher.OutputSlice, selection[random.Int64()])
	}
	return nil
}

func convertToString(cipher *bealeEncryptCipherInfo) (string, error) {
	cipher.OutputText.Grow(len(cipher.OutputSlice))
	for _, i := range cipher.OutputSlice {
		cipher.OutputText.WriteString(fmt.Sprintf("%d, ", i))
	}
	encrypted := cipher.OutputText.String()
	return encrypted[:len(encrypted)-2], nil
}

func collectBealeReferenceFromTxt(key *os.File, cipher *bealeEncryptCipherInfo) error {
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
		err = collectBealeTxtRuneSet(key, cipher)
	case ".pdf":
		cipher.KeyRuneSet, err = collectBealePdfRuneSet(key)
	case ".epub":
		cipher.KeyRuneSet, err = collectBealeEpubRuneSet(key)
	}
	if err != nil {
		return err
	}
	if err = operations.CompareRuneSets(cipher.InputRuneSet, cipher.KeyRuneSet); err != nil {
		return err
	}
	return nil
}

func collectBealeTxtRuneSet(key *os.File, cipher *bealeEncryptCipherInfo) error {
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
