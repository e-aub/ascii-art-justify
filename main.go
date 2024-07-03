package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("tput", "cols")
	cmd.Stdin = os.Stdin
	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(output)
	// left := "\033[5D"
	right := "\033[%dC"
	fmt.Printf(right, string(output))
	fmt.Println("f")
}
