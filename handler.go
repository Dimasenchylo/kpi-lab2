package lab2

import (
	"bytes"
	"io"
)

type ComputeHandler struct {
	input  io.Reader
	output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	inputBuffer, err := io.ReadAll(ch.input)

	if err != nil {
		return err
	}

	trimmedInput := bytes.Trim(inputBuffer, "\x00")

	prefixExpression, err := PostfixToPrefix(string(trimmedInput))
	if err != nil {
		return err
	}

	_, err = ch.output.Write([]byte(prefixExpression))
	if err != nil {
		return err
	}

	return nil
}
