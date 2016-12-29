package main

import (
	"testing"
)

var list = []string{
	"stringA",
	"stringB",
	"stringC",
	"stringD",
	"stringE",
	"stringF",
}

func TestStack(t *testing.T) {
	ps := newProcessingStack(list)
	if len(ps.Loads) != len(list) {
		t.Fatalf("Unexpected size, having %d expecting %d.\n", len(ps.Loads), len(list))
	}
	t.Logf("%#v\n", ps)

	for idx, reference := range list {
		payload, err := ps.GetNext()
		if err != nil {
			t.Fatalf("Unexpected error: %s.\n", err.Error())
		}
		if payload != reference {
			t.Fatalf("Unexpected payload %s != %s at idx %d.\n", payload, reference, idx)
		}
	}
	_, err := ps.GetNext()
	if err == nil {
		t.Fatalf("Should rise the exausted error.\n")
	}

}
