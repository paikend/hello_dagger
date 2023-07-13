package main

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	expected := "Hello, 世界"
	result := HelloWorld()
	if result != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, result)
	}
}
