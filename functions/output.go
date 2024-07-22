package functions

import (
	"errors"
	"fmt"
	"os"
)

// OutputDeliver delivers the output to the console
func OutputDeliver(art string) error {
	if Args.FileName == "" {
		fmt.Print(art)
	} else {
		file, err := os.Create(Args.FileName)
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
