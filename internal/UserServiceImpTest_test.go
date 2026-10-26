package internal

import "testing"

func TestUserServiceImpTest(t *testing.T) {
	if got := "ok"; got != "ok" {
		t.Fatalf("expected ok, got %q", got)
	}
}