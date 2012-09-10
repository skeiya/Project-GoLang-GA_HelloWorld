package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	)

type Gene struct {
	_code string
	_cost int
}

func (g *Gene) Init(length int) {
	buff := make([]byte, length)
	for i:=0; i < length; i++ {
		buff[i]	= byte(rand.Intn(256))
	}
	g._code = string(buff)
	g._cost = 9999
}

func NewGene(length int) *Gene {
	g := new(Gene)
	g.Init(length)
	return g
}

func (g *Gene) CalculateCost(answer string) {
	g._cost = 0
	for i:=0; i<len(g._code); i++ {
		g._cost += int(math.Abs(float64(g._code[i] - answer[i])))
	}
}

type GeneSlice []*Gene

type Population struct {
	_goal string
	_members GeneSlice
}

func NewPopulation(numOfMember int, goal string) *Population {
	p := new(Population)
	p._goal = goal
	p.Init(numOfMember)
	return p
}

func (p *Population) Init(numOfMember int) {
	for i:=0; i < numOfMember; i++ {
		p._members = append(p._members, NewGene(len(p._goal)))
		p._members[i].Init(len(p._goal))
	}
}

// Interface implement for Sort
func (p GeneSlice) Len() int           { return len(p) }
func (p GeneSlice) Less(i, j int) bool { return p[i]._cost < p[j]._cost }
func (p GeneSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func (p *Population) Generation() {
	// calculate cost
	for i:=0; i<len(p._members); i++ {
		p._members[i].CalculateCost(p._goal)
	}

	// sort by cost
	sort.Sort(p._members)

	// if the top score is feasible, finish !
	bestCode := p._members[0]
	fmt.Println(bestCode)
	if (bestCode._cost == 0) {
		return
	}

	// mate the best pair

	// repalace the worst pair with the children

	// mutate
}

func main() {
	population := NewPopulation(20, "Hello world")
	population.Generation()
}
