//go:build integration_tests || all_tests

package all_tests

import "testing"

func TestIntegration(t *testing.T) {
	t.Logf("Ran integration test")
}
