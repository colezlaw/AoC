package registers

import "fmt"

type instruction struct {
	register  *Register
	command   string
	value     int
	Condition *Condition
}

type Instructions struct {
	Series []*instruction
}

func (i *Instructions) Len() int {
	return len(i.Series)
}

func (i *Instructions) Run() {
	for _, ins := range i.Series {
		ins.Execute()
	}
}

func (i *instruction) Execute() error {
	result, err := i.Condition.eval()

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
