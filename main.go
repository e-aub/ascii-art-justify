package main

import (
	"fmt"
	"os"
	"strings"

	f "ascii-art-output/functions"
)

func main() {
	// Check if the provided flags and args are valid
	err := f.ArgsChecker(os.Args[1:])
	f.ErrHandler(err)
	if f.Args.ToDraw == "" {
		return
	}
	f.ToColorIndexes()
	// Minimize the input string
	toMap, err := f.Minimize()
	f.ErrHandler(err)
	// Map the input string to the selected font
	err = f.MapFont(toMap)
	f.ErrHandler(err)
	// Split the input string and Build the output
	spliced := f.Split()

	if (f.Args.FileName != "" && f.Args.ColorFlag.Color != f.Color{}) && (strings.HasSuffix(f.Args.FileName, ".txt")) {
		var Answer string
		fmt.Println(".txt file doesn't support colors, Do you want to continue with Rich Text Format .rtf (y / n)")
		fmt.Scanln(&Answer)
		if Answer == "y" {
			f.Args.FileName = strings.Replace(f.Args.FileName, ".txt", ".rtf", 1)
			fmt.Println(f.Args.FileName)
		}
	}
	art := f.OutputBuilder(spliced)

	// Deliver the output to the console
	err = f.OutputDeliver(art)
	f.ErrHandler(err)
	if f.Args.AlignFlag.Align != "" {
		nSpaces := f.Align()
		fmt.Println(nSpaces)
	}
}
