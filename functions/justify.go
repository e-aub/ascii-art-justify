package functions

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func tput() int {
	command := exec.Command("tput", "cols")
	command.Stdin = os.Stdin
	output, err := command.Output()
	if err != nil {
		ErrHandler(err)
	}
	ttyWidth, err := strconv.Atoi(strings.Trim(string(output), "\n"))
	if err != nil {
		ErrHandler(err)
	}

	return ttyWidth
}

func artWidth() int {
	var width int

	for _, letter := range Args.ToDraw {
		if Args.AlignFlag.Align == "justify" {
			if letter != ' ' {
				width += len(Font[letter][0])
			}
		} else {
			width += len(Font[letter][0])
		}

	}
	return width
}

func Align() int {
	ttyWidth := tput()
	artWidth := artWidth()
	// fmt.Println(ttyWidth)
	// fmt.Println(artWidth)

	switch Args.AlignFlag.Align {
	case "right":
		if margin := ttyWidth - artWidth; margin > 0 {
			return margin
		}
		return 0
	case "center":
		if emptySpace := (ttyWidth - artWidth); emptySpace > 0 {
			return emptySpace / 2
		}
		return 0
	case "justify":
		if wordsNumber := (len(strings.Split(cleanSpaces(), " ")) - 1); wordsNumber > 0 {
			if margin := (ttyWidth - artWidth) / wordsNumber; margin > 0 {
				return margin
			}
		}
		return 0
	}
	return 0
}

func cleanSpaces() string {
	return Spaces.ReplaceAllString(Args.ToDraw, " ")
}
