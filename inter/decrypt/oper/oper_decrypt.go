package oper

import (
	"bufio"
	"os"
	"slices"
	"strconv"
	"strings"
)

// not tested yet
func FileToSlice(input *os.File, inputSlice *[]int, separator string) error {
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

func DecodedSliceToText(outputSlice *[]rune) string {
	var outputText strings.Builder
	outputText.Grow(len(*outputSlice))
	for _, i := range *outputSlice {
		outputText.WriteString(string(i))
	}
	return outputText.String()
}
