package main


type Gene struct {
	_code string
}

func (g *Gene) Init(length int) {
	//_code := "xxx"
}




type Population struct {
	_goal string
	_members []Gene
}

func (p *Population) Init() {
	for i:=0; i < len(p._members); i++ {
		p._members[i].Init(len(p._goal))
	}
}

func (p *Population) Generation()  {
}

func main() {
	population := Population{"hello world", make([]Gene, 20)}
	population.Generation()
}
