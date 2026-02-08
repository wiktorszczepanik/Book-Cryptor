package operations

func ValidateRuneSets(inputSet, keySet map[rune]bool) bool {
	if len(inputSet) >= len(keySet) {
		return false
	}
	for key := range inputSet {
		_, state := keySet[key]
		if !state {
			return false
		}
	}
	return true
}
