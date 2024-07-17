package main

import (
	f "ascii-art-output/functions"
	"fmt"
	"os"
)

func main() {
	// Check if the provided flags and args are valid
	err := f.ArgsChecker(os.Args[1:])
	errHandler(err)
	if f.Arguments.ToDraw == "" {
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
