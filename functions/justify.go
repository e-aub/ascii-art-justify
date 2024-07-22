package functions

import (
	"fmt"
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
		if letter != ' ' {
			width += len(Font[letter][0])

		}
	}
	return width
}

func Align() int {
	ttyWidth := tput()
	artWidth := artWidth()
	fmt.Println(ttyWidth)
	fmt.Println(artWidth)

	switch Args.AlignFlag.Align {
	case "right":
		return ttyWidth - artWidth
	case "center":
		return (ttyWidth - artWidth) / 2
	case "justify":
		return (ttyWidth - artWidth) / (len(strings.Split(cleanSpaces(), " ")) - 1)
	}
	return 0
}

func cleanSpaces() string {
	return Spaces.ReplaceAllString(Args.ToDraw, " ")
}
