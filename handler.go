package lab2

import (
	"bytes"
	"io"
)

type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	inputBuffer, err := io.ReadAll(ch.Input)

	if err != nil {
		return err
	}

	trimmedInput := bytes.Trim(inputBuffer, "\x00")

	prefixExpression, err := PostfixToPrefix(string(trimmedInput))
	if err != nil {
		return err
	}

	_, err = ch.Output.Write([]byte(prefixExpression))
	if err != nil {
		return err
	}

	return nil
}
