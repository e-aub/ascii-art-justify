package functions

import (
	"errors"
	"sort"
	"strings"
)

func Minimize() ([]rune, error) {
	str := strings.ReplaceAll(Args.ToDraw, "\n", "")
	var result []rune
	for _, letter := range str {
		if letter < ' ' || letter > '~' {
			return nil, errors.New("char")
		}
		if !strings.Contains(string(result), string(letter)) {
			result = append(result, letter)
		}
	}
	return sortRunes(result), nil
}

// sort the input string
func sortRunes(runes []rune) []rune {
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return runes
}
