//go:build e2e_tests || all_tests

package all_tests

import "testing"

func TestE2R(t *testing.T) {
	t.Logf("Ran end to end test")
}
