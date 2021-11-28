package utility

import "testing"

func Test_RndNum(t *testing.T) {
	want := 20
	got := RndNum(want)
	if got > want {
		t.Errorf("rndNum(20) = %d, want < %d", got, want)
	}
}
