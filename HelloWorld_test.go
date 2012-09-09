package main

import (
	"testing"
	)

func TestGeneInit(t *testing.T) {
	g := Gene{""}
	str := "Hello World"
	g.Init(len(str))
	if len(g._code) != len(str) {
		t.Error("Gene Init Error:(%d, %d)",  len(g._code), len(str))
	}
	if (g._code == str) {
		t.Error("Random Error")
	}
}
