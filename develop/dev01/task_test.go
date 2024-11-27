package task

import (
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	result := GetTime()

	expected := time.Now()

	if expected.Sub(result) > time.Duration(time.Millisecond*500) {
		t.Errorf("Incorrect result: want %s, have %s", expected, result)
	}
}
