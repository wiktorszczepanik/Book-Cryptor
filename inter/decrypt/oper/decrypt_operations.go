package oper

import (
	"bufio"
	"os"
	"slices"
	"strconv"
	"strings"
)

func FileToSlice(input *os.File, inputSlice *[]int, separator string) error {
	scanner := bufio.NewScanner(input)
	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if i := strings.Index(string(data), separator); i >= 0 {
			return i + len(separator), data[:i], nil
		}
		if atEOF && len(data) > 0 {
			return len(data), data, nil
		}
		return 0, nil, nil
	}
	scanner.Split(split)
	var number int
	var err error
	for scanner.Scan() {
		stringNumber := strings.TrimSpace(scanner.Text())
		if number, err = strconv.Atoi(stringNumber); err != nil {
			return err
		}
		*inputSlice = append(*inputSlice, number)
	}
	input.Seek(0, 0)
	return nil
}

func SortSlice(inputSlice *[]int) []int {
	sorted := make([]int, len(*inputSlice))
	copy(sorted, *inputSlice)
	slices.Sort(sorted)
	return sorted
}

func ReferenceMapToSlice(inputSlice *[]int, keyReferenceMap map[int]rune) *[]rune {
	outputSlice := make([]rune, 0)
	for _, value := range *inputSlice {
		outputSlice = append(outputSlice, keyReferenceMap[value])
	}
	return &outputSlice
}

func DecodedSliceToText(outputSlice *[]rune, inputSlice *[]int, keyReference map[int]rune) string {
	var outputText strings.Builder
	outputText.Grow(len(*outputSlice))
	for _, i := range *inputSlice {
		outputText.WriteString(string(keyReference[i]))
	}
	return outputText.String()
}
