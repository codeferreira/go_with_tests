package dependency

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Jose")

	got := buffer.String()
	want := "Hello, Jose"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
