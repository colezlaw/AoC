package Processing

type Register struct {
	Address string
	Value int
	Owner *Memory
}

func (r *Register) Inc(v int){
	r.Value += v
	r.Owner.ValueChange(r.Value)
}

func (r *Register) Dec(v int){
	r.Value -= v
	r.Owner.ValueChange(r.Value)
}