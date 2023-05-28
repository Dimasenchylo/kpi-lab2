package lab2

import (
	"fmt"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) {
	TestingT(t)
}

type mySuite struct{}

var _ = Suite(&mySuite{})

func (s *mySuite) TestEmpty(c *C) {
	res, err := PostfixToPrefix("")
	c.Assert(res, Equals, "")
	c.Assert(err, ErrorMatches, "invalid expression")
}

func (s *mySuite) TestInvalid(c *C) {
	res, err := PostfixToPrefix("2 +")
	c.Assert(res, Equals, "")
	c.Assert(err, ErrorMatches, "not enough operands")
}

func (s *mySuite) TestSimplePrefixExpression(c *C) {
	res, err := PostfixToPrefix("9 3 /")
	c.Assert(res, Equals, "/ 9 3")
	c.Assert(err, IsNil)
}

func (s *mySuite) TestComplexPrefixExpression(c *C) {
	res, err := PostfixToPrefix("7 12 4 - 2 * - 3 2 1 - ^ 2 * *")
	c.Assert(res, Equals, "* - 7 * - 12 4 2 * ^ 3 - 2 1 2")
	c.Assert(err, IsNil)
}

func ExamplePostfixToPrefix() {
	res, _ := PostfixToPrefix("2 2 +")
	fmt.Println(res)

	// Output:
	// + 2 2
}
