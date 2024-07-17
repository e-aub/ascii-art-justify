package functions

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// OutputBuilder builds the output string
func OutputBuilder(splicedInput []string) string {
	var result strings.Builder
	tracker := 0
	for _, part := range splicedInput {
		if part == "\\n" {
			result.WriteString("\n")
			tracker += 2
			continue
		}
		count := 0
		for count < 8 {
			for i, letter := range part {
				currentIndex := i + tracker
				if InRange(currentIndex) {
					result.WriteString(Arguments.color.color + Font[letter][count] + Colors["reset"])
				} else {
					result.WriteString(Font[letter][count])
				}

			}
			result.WriteString("\n")
			count++
		}
		tracker += len(part) + 2
	}
	return result.String()

}

// OutputDeliver delivers the output to the console
func OutputDeliver(art string) error {
	if !Arguments.output.on {
		fmt.Print(art)
	} else {
		file, err := os.Create(Arguments.output.fileName)
		if err != nil {
			return errors.New("internal")
		}
		_, err = file.WriteString(art)
		defer file.Close()
		if err != nil {
			return errors.New("internal")
		}
	}
	return nil
}
