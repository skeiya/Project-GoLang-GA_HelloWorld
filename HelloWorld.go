package main

import (
	"fmt"
	)

type Gene struct {
	_code string
}

func (g *Gene) Init(length int) {
	//_code := ""
	//for i:=0; i < length; i++ {
	//	_code += "a"
	//}
	g._code = "hello world"
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
	_members []Gene
}

func NewPopulation() *Population {
	p := new(Population)
	p.Init()
	return p
}

func (p *Population) Init() {
	p._members = make([]Gene, 20)
	for i:=0; i < len(p._members); i++ {
		fmt.Print(p._members[i]._code)
		p._members[i].Init(len(p._goal))
	}
}

func (p *Population) Generation()  {
	for i:=0; i<len(p._members); i++ {
		p._members[i].CalculateCost()
	}
}

func main() {
	population := NewPopulation()
	population.Generation()
}
