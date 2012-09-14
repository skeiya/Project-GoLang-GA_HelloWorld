package main

import (
	"testing"
	"sort"
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

	g._code = "aaa"
	g.CalculateCost(answer)
	if (g._cost != 0) {
		t.Error("CalculateCost Error0")
	}
}

func TestSort(t *testing.T) {
	p := NewPopulation(3, "aaa")
	p._members[0]._code = "abb"
	p._members[1]._code = "aab"
	p._members[2]._code = "cab"
	for i:=0; i<len(p._members); i++ {
		p._members[i].CalculateCost("aaa")
	}
	sort.Sort(p._members)
	if (p._members[0]._code != "aab" || p._members[1]._code != "abb") {
		t.Error("Sort Error: " + p._members[0]._code + " <  " + p._members[1]._code)
	}
}

func TestEndCondition(t *testing.T) {
	p := NewPopulation(5, "aaa")
	p._members[0]._code = "abb"
	p._members[1]._code = "caa"
	p._members[2]._code = "ada"
	p._members[3]._code = "aae"
	p._members[4]._code = "aaa"
	p.Generation()
	if (p._members[0]._cost != 0) {
		t.Error("EndCondition Error: " + p._members[0]._code + " <  " + p._members[1]._code + " < " + p._members[2]._code + " < " + p._members[3]._code + " < " + p._members[4]._code)
	}
}

func TestMate(t *testing.T) {
	g1 := NewGene(5)
	g1._code = "aaaaa"
	g2 := NewGene(5)
	g2._code = "bbbbb"

	g1, g2 = g1.Mate(g2)
	if (g1._code != "aabbb" || g2._code != "bbaaa") {
		t.Error("Mate Error")
	}
}
