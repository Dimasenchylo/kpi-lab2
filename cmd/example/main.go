package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	lab2 "github.com/Dimasenchylo/kpi-lab2"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputFromFile   = flag.String("f", "", "Take input from file")
	outputToFile    = flag.String("o", "", "Save output to file")
)

func main() {

	flag.Parse()

	var input io.Reader
	var output io.Writer

	if *inputExpression != "" {
		input = strings.NewReader(*inputExpression)
	} else if *inputFromFile != "" {
		file, err := os.Open(*inputFromFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error opening file:", err)
			os.Exit(1)
		}
		defer file.Close()
		input = file
	}

	if *outputToFile != "" {
		file, err := os.Create(*outputToFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error creating output file:", err)
			os.Exit(1)
		}
		defer file.Close()
		output = file
	} else {
		output = os.Stdout
	}

	if input == nil {
		fmt.Fprintln(os.Stderr, "Input not provided")
		os.Exit(1)
	}

	handler := &lab2.ComputeHandler{
		Input:  input,
		Output: output,
	}

	err := handler.Compute()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error to compute expression:", err)
		os.Exit(1)
	}
}
