package lib1

import "testing"

func Test_outputX_ReturnsCorrectMessage(t *testing.T) {
	want := "Your thing is value\n"
	got := outputX("thing", "value")

	if want != got {
		t.Errorf("outputX, want %s, got %s", want, got)
	}
}
