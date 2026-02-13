package oper

import (
	"bufio"
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"os"
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

func CollectInputSlice(input *os.File, runes *[]rune) error {
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

func ConvertToString(outputSlice *[]int, separator string) (string, error) {
	var outputText strings.Builder
	outputText.Grow(len(*outputSlice))
	for _, i := range *outputSlice {
		outputText.WriteString(fmt.Sprintf("%d"+separator, i))
	}
	encrypted := outputText.String()
	return encrypted[:len(encrypted)-2], nil
}

func SaveOutput(filePath string, cipherText string) error {
	cipherText += "\n"
	var file *os.File
	var err error
	if file, err = os.Create(filePath); err != nil {
		return err
	}
	defer file.Close()
	if _, err = file.WriteString(cipherText); err != nil {
		return err
	}
	return nil
}
