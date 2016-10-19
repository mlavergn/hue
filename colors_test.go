package hue

import (
	"testing"
)

func TestColorMap(t *testing.T) {
	expectXY := [2]float64{0.4432, 0.5154}
	expect := &State{XY: &expectXY, Saturation: 254}
	l1 := ColorYellow(testLights["l1"])
	if l1.XY[0] != expect.XY[0] || l1.XY[1] != expect.XY[1] {
		t.Fatal("State does not match expectation")
	}
}
