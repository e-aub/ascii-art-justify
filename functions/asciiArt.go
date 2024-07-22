package functions

import (
	"fmt"
	"strings"
)

// OutputBuilder builds the output string
func AsciiArt(splicedInput []string) string {
	var result string
	var color string
	var reset string
	rgb := Args.ColorFlag.Color
	if (Args.ColorFlag.Color != Color{} && Args.FileName != "") {
		color = Args.ColorFlag.RtfColor
		reset = Args.ColorFlag.RtfReset
	} else if (Args.ColorFlag.Color != Color{}) && Args.FileName == "" {
		color = fmt.Sprintf(Args.ColorFlag.AnsiColor, rgb.R, rgb.G, rgb.B)
		reset = Args.ColorFlag.AnsiReset
	}
	tracker := 0
	for _, part := range splicedInput {
		if part == "\\n" {
			result += "\n"
			tracker += 2
			continue
		}
		count := 0

		for count < 8 {
			for i, letter := range part {
				margin := Args.AlignFlag.Margin
				currentIndex := i + tracker
				if Args.AlignFlag.Align != "" {
					if (Args.AlignFlag.Align == "right" || Args.AlignFlag.Align == "center") && currentIndex == 0 {
						result += strings.Repeat(" ", margin)
					} else if Args.AlignFlag.Align == "justify" && letter == ' ' {
						result += strings.Repeat(" ", margin)
						continue
					}
				}
				if InRange(currentIndex) {
					result += color + Font[letter][count] + reset
				} else {
					result += Font[letter][count]
				}
			}
			result += "\n"
			count++
		}
		tracker += len(part) + 2
	}
	if (Args.ColorFlag.Color != Color{}) && Args.FileName != "" {
		result = strings.ReplaceAll(result, "\\", "\\\\")
		result = strings.ReplaceAll(result, Args.ColorFlag.RtfReset, Args.ColorFlag.RtfResetCtrlWord)
		result = strings.ReplaceAll(result, Args.ColorFlag.RtfColor, Args.ColorFlag.RtfColorCtrlWord)
		result = strings.ReplaceAll(result, "\n", Args.ColorFlag.NewLineCtrlWord)
		result = fmt.Sprintf(Args.ColorFlag.RtfHeader, rgb.R, rgb.G, rgb.B) + result + "}"
		return result
	}
	return result
}
