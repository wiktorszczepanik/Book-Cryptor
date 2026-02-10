package encrypt

import (
	"book-cryptor/internal/file"
	"book-cryptor/internal/operations"
	"bufio"
	"errors"
	"math/big"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type bealeCipherInfo struct {
	// Files info
	InputFileExt, KeyFileExt   string
	InputFileSize, KeyFileSize int64

	// Data structures
	InputRuneSet, KeyRuneSet map[rune]bool
	KeyReferenceMap          map[rune][]int

	// Output info
	OutputSize  big.Int
	OutputSlice []int64
	OutputText  strings.Builder
}

func EncryptBeale(input, key *os.File) (string, error) {
	cipher := &bealeCipherInfo{}
	var err error
	if err = checkBeale(input, key, cipher); err != nil {
		return "", err
	}
	switch cipher.KeyFileExt {
	case "txt":
		err = encryptBealeFromTxt(input, key, cipher)
	case "pdf":
		err = encryptBealeFromPdf(input, key, cipher)
	case "epub":
		err = encryptBealeFromEpub(input, key, cipher)
	}
	if err != nil {
		return "", err
	}
	var encrypted string
	if encrypted, err = convertToString(cipher); err != nil {
		return "", err
	}
	return encrypted, nil
}

func encryptBealeFromTxt(input, key *os.File, cipher *bealeCipherInfo) error {
	keySize := cipher.KeyFileSize
	inputScanner := bufio.NewScanner(input)
	inputScanner.Split(bufio.ScanWords)
	var err error
	for inputScanner.Scan() {
		word := inputScanner.Text()
		runes := []rune(word)
		for _, r := range runes {
			var runeLocation int64
			if runeLocation, err = findByRuneInTxt(r, key, keySize); err != nil {
				return err
			}
			*&cipher.OutputSlice = append(*&cipher.OutputSlice, runeLocation)
		}
	}
	return nil
}

func findByRuneInTxt(inputRune rune, keyFile *os.File, size int64) (int64, error) {
	// data, err := os.ReadFile(keyFile.Name())
	// if err != nil {
	// 	return 0, err
	// }
	// runes := []rune(string(data))
	// startIndex := rand.Intn(len(runes))
	// for i := 0; i < len(runes); i++ {

	// }

	startOffset := rand.Int63n(size)
	if _, err := keyFile.Seek(startOffset, 0); err != nil {
		return 0, errors.New("Character could not be found in key file.")
	}
	scanner := bufio.NewScanner(keyFile)
	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		character := scanner.Text()
		if character == string(inputRune) {
			return //
		}
	}
}

func encryptBealeFromPdf(input, key *os.File, cipher *bealeCipherInfo) error {
	return nil
}

func encryptBealeFromEpub(input, key *os.File, cipher *bealeCipherInfo) error {
	return nil
}

func convertToString(cipher *bealeCipherInfo) (string, error) {
	cipher.OutputText.Grow(len(cipher.OutputSlice) * 8) // for test
	for i := range *&cipher.OutputSlice {
		cipher.OutputText.WriteString(strconv.Itoa(i) + ", ")
	}
	encrypted := cipher.OutputText.String()
	return encrypted[:len(encrypted)-2], nil
}

func checkBeale(input, key *os.File, cipher *bealeCipherInfo) error {
	var err error
	if *&cipher.InputFileSize, err = file.GetFileSize(input); err != nil {
		return err
	}
	if *&cipher.KeyFileSize, err = file.GetFileSize(key); err != nil {
		return err
	}
	if *&cipher.InputFileExt, err = file.GetInputFileExt(input); err != nil {
		return err
	}
	if *&cipher.KeyFileExt, err = file.GetKeyFileExt(key); err != nil {
		return err
	}
	if *&cipher.InputRuneSet, *&cipher.OutputSize, err = file.CollectTxtRuneSet(input); err != nil {
		return err
	}
	switch *&cipher.KeyFileExt {
	case "txt":
		*&cipher.KeyRuneSet, err = collectBealeTxtRuneSet(key)
	case "pdf":
		*&cipher.KeyRuneSet, err = collectBealePdfRuneSet(key)
	case "epub":
		*&cipher.KeyRuneSet, err = collectBealeEpubRuneSet(key)
	}
	if err != nil {
		return err
	}
	if err = operations.CompareRuneSets(cipher.InputRuneSet, cipher.KeyRuneSet); err != nil {
		return err
	}
	return nil
}

func collectBealeTxtRuneSet(txtFile *os.File) (map[rune]bool, error) {
	runeSet := make(map[rune]bool)
	scanner := bufio.NewScanner(txtFile)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		runeSet[[]rune(word)[0]] = true
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	txtFile.Seek(0, 0)
	return runeSet, nil
}

func collectBealePdfRuneSet(pdfFile *os.File) (map[rune]bool, error) {
	return nil, nil
}

func collectBealeEpubRuneSet(epubFile *os.File) (map[rune]bool, error) {
	return nil, nil
}
