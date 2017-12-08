package Processing

import "fmt"

type Instruction struct {
	register *Register
	command string
	value int
	Condition *Condition
}

type Instructions struct {
	Series []*Instruction
}

func (i *Instructions) Len() int {
	return len(i.Series)
}

func (i *Instructions) Run(){
	for _, instruction := range i.Series {
		instruction.Execute()
	}
}

func (i *Instruction) Execute() error {
	result, err := i.Condition.Eval()

	if err != nil {
		return err
	}

	if result {
		switch i.command {
		case "dec":
			i.register.Dec(i.value)
		case "inc":
			i.register.Inc(i.value)
		default:
			return fmt.Errorf("%v unrecognized command", i.command)
		}
	}

	return nil
}