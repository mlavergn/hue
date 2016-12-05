package hue

import (
  "testing"
)

func TestGetBridge(t *testing.T) {
  br, err := GetBridge()
  if err != nil {
    t.Fatal(err)
  } else if br == nil {
    t.Fatal("Unable to acquire bridge")
  }
}
