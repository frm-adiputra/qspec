package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/frm-adiputra/qspec/internal/gen"
)

const usage = `
Usage:
	qspec SPEC_FILE
`

func main() {
	if len(os.Args) != 2 {
		invalidArgs(errors.New("Invalid number of arguments"))
	}

	err := gen.GenerateSourceCodeFromYAML(os.Args[1])
	if err != nil {
		fatalError(err)
	}
}

func invalidArgs(err error) {
	fmt.Fprintln(os.Stderr, err)
	fmt.Fprintln(os.Stderr, usage)
	os.Exit(1)
}

func fatalError(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
