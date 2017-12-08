package Processing

import "fmt"

type Memory struct {
	registers []*Register
	HighestValue int
}

func (m *Memory) ValueChange(v int){
	if v > m.HighestValue {
		m.HighestValue = v
	}
}

func (m *Memory) FindLargestValue() int {
	max := 0
	for _, r := range m.registers {
		if r.Value > max {
			max = r.Value
		}
	}
	return max
}

func (m *Memory) Len() int {
	return len(m.registers)
}

func (m *Memory) Add(address string) *Register{
	r,err := m.Find(address)
	if err == nil {
		return r
	}
	r = &Register{address, 0, m}
	m.registers = append(m.registers, r)
	return r
}

func (m Memory) Find(address string) (*Register, error) {
	for _, r := range m.registers{
		if r.Address == address {
			return r, nil
		}
	}

	return nil, fmt.Errorf("address not found")
}
