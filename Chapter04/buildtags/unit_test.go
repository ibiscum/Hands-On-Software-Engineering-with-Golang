//go:build unit_tests || all_tests

package all_tests

import "testing"

func TestUnit(t *testing.T) {
	t.Logf("Ran unit test")
}
