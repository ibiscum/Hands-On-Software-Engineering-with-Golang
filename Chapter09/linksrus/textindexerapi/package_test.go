package textindexerapi_test

import (
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/protobuf/types/known/timestamppb"

	gc "gopkg.in/check.v1"
)

var err error

func Test(t *testing.T) {
	// Run all gocheck test-suites
	// gc.TestingT(t)
}

func mustEncodeTimestamp(c *gc.C, t time.Time) *timestamp.Timestamp {
	ts := timestamppb.New(t)
	c.Assert(err, gc.IsNil)
	return ts
}
