package operations

import (
	"io"
	"strings"
	"testing"

	"github.com/go-check/check"
)

func Test(t *testing.T) { check.TestingT(t) }

type CheckSuite struct{}

var _ = check.Suite(&CheckSuite{})

func (s *CheckSuite) TestCheckSimple(c *check.C) {
	c.Assert(nil, check.IsNil)
	c.Check(42, check.Equals, 42)
	c.Assert(strings.ToLower("HELLO"), check.Equals, "hello")
	c.Assert(io.ErrClosedPipe, check.ErrorMatches, "io: .* on closed pipe")
}
