package main

import(
	"io/ioutil"
	"strings"
	"fmt"
	"strconv"
)

type Program struct {
	name string
	weight int
	aggregate int
	namesSupported []string
	base *Program
	tower []*Program
}

type Mason struct {
	programs []*Program
	cornerstone *Program
}

func (p *Program) aggregateWeight() int {
	if p.aggregate > 0 {
		return p.aggregate
	}

	p.aggregate = p.weight

	if p.spire() {
		return p.aggregate
	}

	for _, program := range p.tower {
		p.aggregate += program.aggregateWeight()
	}

	return p.aggregate

}

func (p *Program) spire() bool  {
	return len(p.namesSupported) < 1
}

func (m *Mason) build () {
	for _, p := range m.programs{
		if p.spire() {
			continue
		}
		for _, name := range p.namesSupported{
			program := m.find(name)
			program.base = p
			p.tower = append(p.tower, program)
		}
	}
	setCornerstone(m)

}


func setCornerstone(m *Mason) {
	for _, p := range m.programs {
		if p.base != nil {
			continue
		}

		m.cornerstone = p
		break
	}
}

func (m *Mason) find(name string) *Program {
	for _, p := range m.programs{
		if p.name == name {
			return p
		}
	}
	return nil
}

func (m *Mason) materials (s []string)  {
	for _, line := range s{
		var p Program
		p.init(strings.Fields(line))
		m.collect(&p)
	}
}

func (m *Mason) collect (p *Program){
	m.programs = append(m.programs, p)
}

func (p *Program) init (f []string)  {
	p.name = f[0]
	p.weight = parseWeight(f[1])

	if len(f) > 3 {
		names := f[3:len(f)]
		for k,v := range names{
			names[k] = strings.Trim(v, ",")
		}
		p.namesSupported = names
	}
}

func (p *Program) balanced() bool {
	if p.spire() {
		return true
	}

	unique := uniqueWeight(p)

	if len(unique) > 1 {
		return false
	}

	return true
}
func uniqueWeight(p *Program) map[int]bool {
	unique := map[int]bool{}
	for _, supported := range p.tower {
		unique[supported.aggregateWeight()] = true
	}
	return unique

}

func parseWeight(w string) int {
	weight, _ := strconv.Atoi(strings.Trim(w, "()"))
	return  weight
}

func input() []string {
	content, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n")
	return lines
}

func (p *Program) lastUnbalanced() *Program {
	var last *Program
	next := p.nextUnbalanced()
	for true {
		if next != nil {
			last = next
			next = last.nextUnbalanced()
		}else{
			break
		}
	}
	return last
}

func (p *Program) nextUnbalanced() *Program {
	for _, program := range p.tower {
		if program.balanced() {
			continue
		}
		return program
	}
	return nil
}

func (m *Mason) balance() (int, *Program) {
	var unbalanced *Program
	var expected, unexpected int
	weights := map[int]int{}
	last := m.cornerstone.lastUnbalanced()

	for _,p := range last.tower {
		weights[p.aggregateWeight()] +=1
	}

	for k,v := range weights {
		if v > 1 {
			expected = k
		}else{
			unexpected = k
		}
	}

	for _,p := range last.tower {
		if p.aggregateWeight() == unexpected {
			unbalanced = p
			//adjust = p.weight
			break
		}

	}

	return unbalanced.weight + (expected-unexpected), unbalanced
}


func main()  {
	var mason Mason
	mason.materials(input())
	mason.build()
	fmt.Printf("%v is the cornerstone of the towers.\n", mason.cornerstone.name)
	newWeight, unbalancedProgram := mason.balance()
	fmt.Printf("%v currently weighs %v and needs to weigh %v\n", unbalancedProgram.name, unbalancedProgram.weight, newWeight )

}
