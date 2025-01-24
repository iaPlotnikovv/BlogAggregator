package app

import (
	"fmt"
	"os"
)

// errors.go
func handleError(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
