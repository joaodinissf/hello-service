package reverse

import "testing"

func TestStringReversesAName(t *testing.T) {
	if got := String("Drive"); got != "evirD" {
		t.Errorf("String(%q) = %q, want %q", "Drive", got, "evirD")
	}
}
