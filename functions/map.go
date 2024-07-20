package functions

import (
	"bufio"
	"errors"
	"os"
)

var Font map[rune][]string

func MapFont(toMap []rune) error {
	Font = make(map[rune][]string)
	file, err := os.Open("./templates/" + Args.Banner)
	if err != nil {
		return errors.New("fs")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lineCount := 0
	rIndex := 0
	for scanner.Scan() {
		if lineCount == ((int(toMap[rIndex])-31)*9)-9 {
			i := 0
			for i < 8 && scanner.Scan() {
				Font[toMap[rIndex]] = append(Font[toMap[rIndex]], scanner.Text())
				lineCount++
				i++
			}
			if rIndex != len(toMap)-1 {
				rIndex++
			}
		}
		lineCount++
	}
	return nil
}
