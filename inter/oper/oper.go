package oper

import (
	"bufio"
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"os"
	"slices"
	"strconv"
	"strings"
)

func CompareRuneSets(inputSet, keySet map[rune]bool) error {
	if len(inputSet) > len(keySet) {
		return errors.New("Not enough available key file characters to encrypt input file.")
	}
	for key := range inputSet {
		_, state := keySet[key]
		if !state {
			return errors.New("Not enough available key file characters to encrypt input file.")
		}
	}
	return nil
}

func CollectPlainInputSlice(input *os.File, runes *[]rune) error {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		for _, r := range word {
			*runes = append(*runes, r)
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func GenerateCipher(inputSlice []rune, keyReferenceMap map[rune][]int) ([]int, error) {
	outputSlice := make([]int, 0, len(inputSlice))
	for _, v := range inputSlice {
		selection := keyReferenceMap[v]
		selectionLength := int64(len(selection))
		random, err := rand.Int(rand.Reader, big.NewInt(selectionLength))
		if err != nil {
			return nil, err
		}
		outputSlice = append(outputSlice, selection[random.Int64()])
	}
	return outputSlice, nil
}

func ConvertEncryptedSliceToString(outputSlice *[]int, separator string) (string, error) {
	var outputText strings.Builder
	outputText.Grow(len(*outputSlice))
	for _, i := range *outputSlice {
		outputText.WriteString(fmt.Sprintf("%d"+separator, i))
	}
	encrypted := outputText.String()
	return encrypted[:len(encrypted)-2], nil
}

func CollectAllPlainTxtRuneSet(txtFile *os.File) (map[rune]bool, int, error) {
	runeSet := make(map[rune]bool)
	scanner := bufio.NewScanner(txtFile)
	scanner.Split(bufio.ScanWords)
	var sizeCouner int = 0
	for scanner.Scan() {
		word := scanner.Text()
		for _, letter := range word {
			runeSet[letter] = true
			sizeCouner++
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, 0, err
	}
	txtFile.Seek(0, 0)
	return runeSet, 0, nil
}

func CollectPdfRuneSet(txtFile *os.File) map[rune]bool {
	return nil
}

func CollectEpubRuneSet(txtFile *os.File) map[rune]bool {
	return nil
}

// not tested yet
func EncryptedFileToSlice(input *os.File, inputSlice *[]int, separator string) error {
	scanner := bufio.NewScanner(input)
	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		sepLength := len(separator)
		for i := 0; i < len(data); i++ {
			if string(data[i:i+sepLength]) == separator {
				return i + 1, data[:i], nil
			}
		}
		if atEOF {
			return 0, nil, nil
		}
		return 0, data, bufio.ErrFinalToken
	}
	scanner.Split(split)
	var number int
	var err error
	for scanner.Scan() {
		if number, err = strconv.Atoi(scanner.Text()); err != nil {
			return err
		}
		*inputSlice = append(*inputSlice, number)
	}
	return nil
}

func GetSortedEncryptedInputSlice(inputSlice *[]int) []int {
	sorted := make([]int, len(*inputSlice))
	copy(sorted, *inputSlice)
	slices.Sort(sorted)
	return sorted
}

func ConvertReferenceMapToSlice(inputSlice *[]int, keyReferenceMap map[int]rune) *[]rune {
	outputSlice := make([]rune, 0)
	for _, value := range *inputSlice {
		outputSlice = append(outputSlice, keyReferenceMap[value]) 
	}
	return &outputSlice
} 

func ConvertDecodedSliceToText(outputSlice *[]rune) string {
	var outputText strings.Builder
	outputText.Grow(len(*outputSlice))
	for _, i := range *outputSlice {
		outputText.WriteString(string(i))
	}
	return outputText.String()
}