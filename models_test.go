package hue

import (
	"testing"
)

func TestModel(t *testing.T) {
	expect := "Hue A19"
	l1 := Model(testLights["l1"])
	if l1 != expect {
		t.Fatal("expected: " + expect + ", got: " + l1)
	}
}

func TestGamut(t *testing.T) {
	expect := GamutA
	l1 := Gamut(testLights["l1"])
	if l1 != expect {
		t.Fatal("expected: " + expect + ", got: " + l1)
	}
}
