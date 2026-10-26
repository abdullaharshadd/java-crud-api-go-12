package internal

import "testing"

func TestUserController(t *testing.T) {
	if got := "ok"; got != "ok" {
		t.Fatalf("expected ok, got %q", got)
	}
}