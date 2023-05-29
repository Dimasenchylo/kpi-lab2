package lab2

import (
	"bytes"
	. "gopkg.in/check.v1"
	"strings"
)

type ComputeHandlerSuite struct{}

var _ = Suite(&ComputeHandlerSuite{})

func (s *ComputeHandlerSuite) Test1(c *C) {
	input := strings.NewReader("5 2 1 - +")
	output := bytes.NewBuffer(nil)

	ch := &ComputeHandler{
		input:  input,
		output: output,
	}

	err := ch.Compute()

	c.Check(err, IsNil)
	c.Check(output.String(), Equals, "+ 5 - 2 1")

}

func (s *ComputeHandlerSuite) Test2(c *C) {
	input := strings.NewReader("9 + $ ? 7 - 3")
	output := bytes.NewBuffer(nil)

	ch := &ComputeHandler{
		input:  input,
		output: output,
	}

	err := ch.Compute()

	c.Check(err, NotNil)
}
