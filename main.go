package main

import (
	"adiazsor/aluminum/util"
	"fmt"
	"os"
	"strings"
)

var inputFile string

func init() {
	if len(os.Args) >= 2 {
		inp := strings.TrimSpace(os.Args[1])
		err := util.ValidateInputFile(inp)
		if err != nil {
			fmt.Printf("Issue with input file: %s\n", err)
			util.PrintUsageStdout()
			os.Exit(1)
		}

		inputFile = inp
	} else {
		util.PrintUsageStdout()
	}
}

func main() {
	fmt.Println(inputFile)
}
