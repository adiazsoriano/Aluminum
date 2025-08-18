package util

import (
	"fmt"
	"os"
	"strings"
)

func PrintUsageStdout() {
	fmt.Printf("Usage: %s [inputfile.alum] ...\n", os.Args[0])
}

func ValidateInputFile(inp string) error {
	inpParts := strings.Split(inp, ".")
	if len(inpParts) <= 1 {
		return fmt.Errorf("make sure that the file \"%s\" ends in \".alum\"", inp)
	}

	if inpParts[len(inpParts)-1] != "alum" {
		return fmt.Errorf("make sure that the file \"%s\" contains and ends with \".alum\"", inp)
	}

	return nil
}
