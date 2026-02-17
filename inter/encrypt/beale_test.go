package encrypt

import (
	"book-cryptor/inter/file"
	"os"
	"strings"
	"testing"
)

var inputFile *os.File
var keyFile *os.File
var err error

func TestBealePlainTextLoremIpsum(t *testing.T) {
	const inputFilePath = "../../test/txt/SmallLoremIpsum.txt"
	const keyFilePath = "../../test/txt/LoremIpsum.txt"
	const separator = ", "
	if inputFile, keyFile, err = file.GetEssensialFiles(inputFilePath, keyFilePath); err != nil {
		t.Errorf(`Can't get essensial files`, inputFile, keyFile, err)
	}
	// plaintext, err := file.FileContentToString(inputFile)
	// if err != nil {
	// 	t.Errorf(`Can't read input file`, inputFile, keyFile, err)
	// }

	// Example "want"
	// 1, 63, 42, 11, 65, 67, 52, 54, 28, 18
	want := [][]string{
		{"1"},        // L: 1
		{"55", "63"}, // o: 2
		{"42"},       // r: 1
		{"8", "11", "16", "21", "27", "33", "34", "46", "49", "68"}, // e: 10
		{"18", "23", "65"},                        // m: 3
		{"2", "13", "39", "41", "43", "60", "67"}, // i: 7
		{"52", "58"},                              // p: 2
		{"4", "9", "54", "59"},                    // s: 4
		{"14", "28", "41"},                        // u: 3
		{"18", "23", "65"},                        // m: 3
	}
	msg, err := Beale(inputFile, keyFile, separator)
	if err != nil {
		t.Errorf(`%q, %v, want match for %#q, nil`, msg, err, want)
	}
	msgSlice := strings.Split(msg, separator)
	for i := 0; i < len(want); i++ {
		var found bool = false
		for j := 0; j < len(want[i]); j++ {
			if plaintext[j] == want[i][j] {
				found = true
			}
		}
		if !found {
			t.Errorf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
		}
	}
}
