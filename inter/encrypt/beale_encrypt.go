package encrypt

import (
	"book-cryptor/inter/encrypt/oper"
	"book-cryptor/inter/file"
	"bufio"
	"errors"
	"os"
	"strings"
	"unicode"
)

type bealeEncryption interface {
	collectBealeTxtRuneSet(key *os.File, exact bool) error
	collectBealeReferenceMapFromTxt(file *os.File, exact bool) error
}

type bealeEncryptInfo struct {
	// Files info
	KeyFileExt string

	// Data structures
	InputSlice               []rune
	InputRuneSet, KeyRuneSet map[rune]bool
	KeyReferenceMap          map[rune][]int

	// Output info
	OutputSize  int
	OutputSlice []int
	OutputText  strings.Builder
}

func Beale(input, key *os.File, separator string, exact bool) (string, error) {
	var bealeCipher bealeEncryption = &bealeEncryptInfo{}
	var err error
	var output string = ""
	if cipher, ok := bealeCipher.(*bealeEncryptInfo); ok {
		if err = checkBeale(input, key, cipher, exact); err != nil {
			return "", err
		}
		err = oper.CollectPlainSlice(input, &cipher.InputSlice, exact)
		if err != nil {
			return "", nil
		}
		switch cipher.KeyFileExt {
		case ".txt":
			err = cipher.collectBealeReferenceMapFromTxt(key, exact)
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
		if output, err = oper.ConvertSliceToString(&cipher.OutputSlice, separator); err != nil {
			return "", err
		}
	} else {
		return "", errors.New("Invalid assertion type in beale operations.")
	}
	return output, nil
}

func checkBeale(input, key *os.File, cipher *bealeEncryptInfo, exact bool) error {
	var err error
	if err = file.CheckInputFileExt(input); err != nil {
		return err
	}
	if cipher.KeyFileExt, err = file.GetKeyFileExt(key); err != nil {
		return err
	}
	cipher.InputRuneSet, cipher.OutputSize, err = oper.CollectPlainTxtRuneSet(input, exact)
	if err != nil {
		return err
	}
	switch cipher.KeyFileExt {
	case ".txt":
		err = cipher.collectBealeTxtRuneSet(key, exact)
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

func (cipher *bealeEncryptInfo) collectBealeTxtRuneSet(key *os.File, exact bool) error {
	cipher.KeyRuneSet = make(map[rune]bool)
	scanner := bufio.NewScanner(key)
	scanner.Split(bufio.ScanWords)
	if exact {
		for scanner.Scan() {
			word := scanner.Text()
			cipher.KeyRuneSet[[]rune(word)[0]] = true
		}
	} else {
		for scanner.Scan() {
			word := scanner.Text()
			if unicode.IsLetter([]rune(word)[0]) {
				cipher.KeyRuneSet[unicode.ToLower([]rune(word)[0])] = true
			}
		}
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

func (cipher *bealeEncryptInfo) collectBealeReferenceMapFromTxt(key *os.File, exact bool) error {
	cipher.KeyReferenceMap = make(map[rune][]int)
	keyScanner := bufio.NewScanner(key)
	keyScanner.Split(bufio.ScanWords)
	var runeCounter int = 1
	if exact {
		for keyScanner.Scan() {
			firstRune := ([]rune(keyScanner.Text()))[0]
			if cipher.InputRuneSet[firstRune] {
				cipher.KeyReferenceMap[firstRune] = append(cipher.KeyReferenceMap[firstRune], runeCounter)
			}
			runeCounter++
		}
	} else {
		for keyScanner.Scan() {
			var firstRune rune
			if unicode.IsLetter(([]rune(keyScanner.Text()))[0]) {
				firstRune = unicode.ToLower(([]rune(keyScanner.Text()))[0])
				if cipher.InputRuneSet[firstRune] {
					cipher.KeyReferenceMap[firstRune] = append(cipher.KeyReferenceMap[firstRune], runeCounter)
				}
			}
			runeCounter++
		}
	}
	if err := keyScanner.Err(); err != nil {
		return err
	}
	key.Seek(0, 0)
	return nil
}

func encryptBealeFromPdf(input, key *os.File, cipher *bealeEncryptInfo) error {
	return nil
}

func encryptBealeFromEpub(input, key *os.File, cipher *bealeEncryptInfo) error {
	return nil
}
