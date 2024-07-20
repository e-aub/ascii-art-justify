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
	errHandler(err)
	if f.Args.ToDraw == "" {
		return
	}
	f.ToColorIndexes()
	// Minimize the input string
	toMap, err := f.Minimize()
	errHandler(err)
	// Map the input string to the selected font
	err = f.MapFont(toMap)
	errHandler(err)
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
	errHandler(err)
}

func errHandler(err error) {
	if err != nil {
		fmt.Println(f.Errors[err.Error()])
		os.Exit(1)
	}
}
