package main

import (
	"fmt"
	"math/rand"
	)

type Gene struct {
	_code string
}

func (g *Gene) Init(length int) {
	buff := make([]byte, length)
	for i:=0; i < length; i++ {
		buff[i]	= byte(rand.Intn(256))
	}
	g._code = string(buff)
}

func NewGene(length int) *Gene {
	g := new(Gene)
	g.Init(length)
	return g
}

func (g *Gene) CalculateCost() {
	for i:=0; i<len(g._code); i++ {
		fmt.Println(g._code[i])
	}
}


type Population struct {
	_goal string
	_members [](*Gene)
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

func (p *Population) Generation()  {
	for i:=0; i<len(p._members); i++ {
		p._members[i].CalculateCost()
	}
}

func main() {
	population := NewPopulation(20, "Hello world")
	population.Generation()
}
