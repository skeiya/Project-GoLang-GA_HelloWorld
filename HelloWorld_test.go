package main

import (
	"testing"
	)

func TestGeneInit(t *testing.T) {
	str := "Hello World"
	g := NewGene(len(str))
	if len(g._code) != len(str) {
		t.Error("Gene Init Error:(%d, %d)",  len(g._code), len(str))
	}
	if (g._code == str) {
		t.Error("Random Error")
	}
}

func TestPopulationInit(t *testing.T) {
	numOfMembers := 20
	p := NewPopulation(numOfMembers, "Hello world")
	if (len(p._members) != numOfMembers) {
		t.Error("Population Init Error")
	}
}

func TestGeneCost(t *testing.T) {
	answer := "aaa"
	g := NewGene(len(answer))
	g._code = "aab"
	g.CalculateCost(answer)
	if (g._cost != 1) {
		t.Error("CalculateCost Error")
	}
}

func TestSort(t *testing.T) {
	p := NewPopulation(2, "aaa")
	p._members[0]._code = "abb"
	p._members[1]._code = "aab"
	p.Generation()
	if (p._members[0]._code != "aab" || p._members[1]._code != "abb") {
		t.Error("Sort Error")
	}
}

func TestEndCondition(t *testing.T) {
	p := NewPopulation(2, "aaa")
	p._members[0]._code = "abb"
	p._members[1]._code = "aaa"
	p.Generation()
	if (p._members[0]._cost != 0) {
		t.Error("EndCondition Error")
	}

}
