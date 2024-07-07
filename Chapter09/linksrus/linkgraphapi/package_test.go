package linkgraphapi_test

import (
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/protobuf/types/known/timestamppb"
	gc "gopkg.in/check.v1"
)

func Test(t *testing.T) {
	// Run all gocheck test-suites
	gc.TestingT(t)
}

func mustEncodeTimestamp(t time.Time) *timestamp.Timestamp {
	ts := timestamppb.New(t)
	return ts
}

func mustDecodeTimestamp(c *gc.C, ts *timestamp.Timestamp) time.Time {
	t, err := ptypes.Timestamp(ts)
	c.Assert(err, gc.IsNil)
	return t
}
