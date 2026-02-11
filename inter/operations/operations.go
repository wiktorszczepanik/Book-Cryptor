package operations

import (
	"errors"
)

var ErrMissingChars = errors.New("Not enough available key file characters to encrypt input file.")

func CompareRuneSets(inputSet, keySet map[rune]bool) error {
	if len(inputSet) > len(keySet) {
		return ErrMissingChars
	}
	for key := range inputSet {
		_, state := keySet[key]
		if !state {
			return ErrMissingChars
		}
	}
	return nil
}
