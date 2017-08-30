package main

import (
	"testing"
)

func TestX(t *testing.T) {
	if true != true {
		t.Fatalf("We have a problem. Got %v, want %v", true, true)
	}
}
