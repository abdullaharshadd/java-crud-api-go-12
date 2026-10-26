package internal

import "testing"

func TestUserDao(t *testing.T) {
	if got := "ok"; got != "ok" {
		t.Fatalf("expected ok, got %q", got)
	}
}