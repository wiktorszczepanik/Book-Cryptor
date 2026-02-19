package encrypt

import (
	"book-cryptor/inter/file"
	"os"
	"strings"
	"testing"
)

func TestExactBealePlainTextLoremIpsum(t *testing.T) {
	const smallLoremIpsum = "Loremipsum"
	const inputFilePath = "test/PlainLoremIpsum.txt"
	const keyFilePath = "test/KeyLoremIpsum.txt"
	const separator = ", "
	const exact = false
	want := [][]string{
		{"1"},        // L: 1
		{"55", "63"}, // o: 2
		{"42"},       // r: 1
		{"8", "11", "16", "21", "27", "33", "34", "46", "49", "68"}, // e: 10
		{"18", "23", "65"},                        // m: 3
		{"2", "13", "39", "41", "43", "60", "67"}, // i: 7
		{"52", "58"},                              // p: 2
		{"4", "9", "54", "59"},                    // s: 4
		{"14", "28", "31"},                        // u: 3
		{"18", "23", "65"},                        // m: 3
	}
	var inputFile, keyFile *os.File
	var err error
	if inputFile, keyFile, err = file.GetEssensialFiles(inputFilePath, keyFilePath); err != nil {
		t.Errorf(`Can't read essensial files: %v, %v, %v`, inputFile, keyFile, err)
	}
	defer inputFile.Close()
	defer keyFile.Close()
	plaintext, err := file.FileContentToString(inputFile)
	if err != nil {
		t.Errorf(`Can't read input file: %v, %v, %v`, inputFile, keyFile, err)
	}
	if plaintext != smallLoremIpsum {
		t.Errorf(`%v, want match for %v`, plaintext, smallLoremIpsum)
	}
	msg, err := Beale(inputFile, keyFile, separator, exact)
	if err != nil {
		t.Errorf(`%v, Beale encryption - %v`, msg, err)
	}
	msgSlice := strings.Split(msg, separator)
	for i := 0; i < len(want); i++ {
		var found bool = false
		for j := 0; j < len(want[i]); j++ {
			if msgSlice[i] == want[i][j] {
				found = true
			}
		}
		if !found {
			t.Errorf(`%v, %v, want match for %v, nil`, msgSlice[i], err, want[i])
		}
	}
}
