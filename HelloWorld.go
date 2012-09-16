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

func GetRandChar() string {
	for {
		c := rand.Intn(256)
		if 0x20 < c && c < 0x7f {
			return string(c)
		}	
	}
	return "a"
}

func (g *Gene) Init(length int) {
	str := ""
	for i:=0; i < length; i++ {
		str += GetRandChar()
	}
	g._code = str
	g._cost = 9999
}

func NewGene(length int) *Gene {
	g := new(Gene)
	g.Init(length)
	return g
}

func Mutate(org uint8) string {
	isInc := 1
	if (0.5 < rand.Float32()) {
		isInc = -1
	}
	org += uint8(isInc)
	if (org == 0x7f) {
		org = 0x20
	} else if (org == 0x1f) {
		org = 0x7f
	}
	return string(org)
}

func (g *Gene) Mutate(chance float32) {
	if (chance < rand.Float32()) {
		return
	}
	
	length := len(g._code)
	index := rand.Intn(length)
	org := g._code
	g._code = ""
	for i:=0; i<length; i++ {
		if (i == index) {
			g._code += string(Mutate(org[i]))
		} else {
			g._code += string(org[i])
		}
	}
}

func (g *Gene) CalculateCost(answer string) {
	g._cost = 0
	for i:=0; i<len(g._code); i++ {
		g._cost += int(math.Abs(float64(g._code[i]) - float64(answer[i])))
	}
}

func (g1 *Gene) Mate(g2 *Gene) (*Gene, *Gene) {
	length := len(g1._code)
	new_g1 := NewGene(length)
	new_g2 := NewGene(length)

	mid := length/2
	new_g1._code = g1._code[0:mid] + g2._code[mid:length]
	new_g2._code = g2._code[0:mid] + g1._code[mid:length]
	return new_g1, new_g2
}

type GeneSlice []*Gene

type Population struct {
	_generation int
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
	p._generation = 0
	for i:=0; i < numOfMember; i++ {
		p._members = append(p._members, NewGene(len(p._goal)))
		p._members[i].Init(len(p._goal))
	}
}

// Interface implement for Sort
func (p GeneSlice) Len() int           { return len(p) }
func (p GeneSlice) Less(i, j int) bool { return p[i]._cost < p[j]._cost }
func (p GeneSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func (p *Population) Dump() {
	for i:=2; i<len(p._members); i++ {
		msg := fmt.Sprintf("(%d)", p._members[i]._cost)
		fmt.Print(p._members[i]._code + msg + "   ")
	}
	fmt.Println("\n------------------------------------------------------------------\n")
}

func (p *Population) Generation() {
	// calculate cost
	for i:=0; i<len(p._members); i++ {
		p._members[i].CalculateCost(p._goal)
	}

	// sort by cost
	sort.Sort(p._members)

	// if the top score is feasible, finish !
	bestCode := p._members[0]
	fmt.Print("top:" + p._members[0]._code + "," + p._members[1]._code)
	fmt.Printf(" (%d, %d), generation=%d\n",p._members[0]._cost, p._members[1]._cost, p._generation)
	if (bestCode._cost == 0) {
		return
	}

	// mate the best pair
	g1, g2 := p._members[0].Mate(p._members[1])

	// repalace the worst pair with the children
	p._members = append(p._members[0:len(p._members)-2], g1, g2)

	// mutate
	for i:=2; i<len(p._members)-2; i++ {
		p._members[i].Mutate(0.5)
	}

	p._generation++
	p.Generation()
}

func main() {
	population := NewPopulation(20, "Hello world")
	population.Generation()
}
