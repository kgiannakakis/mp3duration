package mp3duration

import "testing"

func TestCbr(t *testing.T) {
	duration, _ := Calculate("testdata/demo - cbr.mp3")

	if duration != 285.78 {
		t.Errorf("FAILED: Expected 285.78 got %v", duration)
	} else {
		t.Log("PASSED")
	}
}

func TestVbr(t *testing.T) {
	duration, _ := Calculate("testdata/demo - vbr.mp3")

	if duration != 285.727 {
		t.Errorf("FAILED: Expected 285.727 got %v", duration)
	} else {
		t.Log("PASSED")
	}
}
