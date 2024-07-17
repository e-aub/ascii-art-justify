package main

import (
	f "ascii-art-output/functions"
	"fmt"
	"log"
	"os"
	"testing"
)

func mainCopy(args []string) string {
	f.Params.Args = args
	f.FlagChecker()
	if f.Params.Err != nil {
		return f.Params.Err.Error()
	}
	f.ArgsChecker()
	if f.Params.Err != nil {
		fmt.Println()
		return f.Params.Err.Error()
	}

	if f.Params.Input == "" {
		return ""
	}

	f.Minimize()
	if f.Params.Err != nil {
		return f.Params.Err.Error()
	}

	f.MapFont()
	if f.Params.Err != nil {
		return f.Params.Err.Error()
	}

	f.Split()
	f.OutputBuilder()
	if f.Params.Err != nil {
		return f.Params.Err.Error()
	}

	if f.Params.OutputFile != "" {
		f.OutputDeliver()
		if f.Params.Err != nil {
			return f.Params.Err.Error()
		}
	} else {
		return f.Params.Output
	}

	return ""
}

func TestMain(t *testing.T) {
	flags := []string{"", "--output=test.txt", "-output=test.txt", "--output=test", "--output"}
	tests := []string{"0123456789", "ABCDEFGHIJKLMNOPQRSTUVWXYZ", "abcdefghijklmnopqrstuvwxyz", "/(\")", "hello1\\nworld"}
	banners := []string{"", "standard", "shadow", "thinkertoy", "enigma", "nirvana"}
	fileName := "test_files/want%d%s.txt"
	var want string
	for j, flag := range flags {
		for i, test := range tests {
			for _, banner := range banners {
				args := []string{}
				if flag != "" {
					args = append(args, flag)
				}
				args = append(args, test)
				if banner != "" {
					args = append(args, banner)
				} else {
					banner = "standard"

				}
				got := mainCopy(args)
				if j == 1 {
					got = readFile("test.txt")
				}

				want = readFile(fmt.Sprintf(fileName, i+1, banner))
				if j > 1 {
					want = "Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard"
				}

				if string(got) != want {
					t.Errorf("Test case %d failed. Expected: %q", i+1, fmt.Sprintf(fileName, i+1, banner))
				}
			}
		}

	}
	os.Remove("test.txt")
}

func readFile(fileName string) string {
	content, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}
