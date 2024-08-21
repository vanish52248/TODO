package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat(3, "a")
	expected := "aaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}
