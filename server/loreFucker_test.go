package main

import (
	"testing"
)

func TestGeneration(t *testing.T) {
	value, err := GenerateSubnet("127.0.0.0", 4, "127.0.10.0")
	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	t.Log(value)
	value, err = GenerateSubnet("127.0.0.0", 4, "127.1.10.0")
	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	t.Log(value)

}
