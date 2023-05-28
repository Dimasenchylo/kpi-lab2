package lab2

import (
	"fmt"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t testing.T) {
	TestingT(t)
}

type mySuite struct{}

var _ = Suite(&mySuite{})

func (smySuite) TestEmpty(c C) {
	res, err := PostfixToPrefix("")
	c.Assert(res, Equals, "")
	c.Assert(err, ErrorMatches, "invalid expression")
}

func (smySuite) TestInvalid(c C) {
	res, err := PostfixToPrefix("2 +")
	c.Assert(res, Equals, "")
	c.Assert(err, ErrorMatches, "not enough operands")
}

func (smySuite) TestSimplePrefixExpression(c C) {
	res, err := PostfixToPrefix("9 3 /")
	c.Assert(res, Equals, "/ 9 3")
	c.Assert(err, IsNil)
}

func (smySuite) TestComplexPrefixExpression(c C) {
	res, err := PostfixToPrefix("1 9 + 7 3 - 5 2 ^ + 8 4 / - 6 + ")
	c.Assert(res, Equals, " + 1 - 9 + 7 3 - ^ 5 2 + / 8 4 6")
	c.Assert(err, IsNil)
}

func ExamplePostfixToPrefix() {
	res, _ := PostfixToPrefix("2 2 +")
	fmt.Println(res)

	// Output:
	// + 2 2
}
