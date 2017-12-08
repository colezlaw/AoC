package registers

import "fmt"

type Condition struct {
	left     *Register
	operator string
	right    int
}

func (c *Condition) eval() (bool, error) {
	switch c.operator {

	case ">":
		return c.left.Value > c.right, nil
	case ">=":
		return c.left.Value >= c.right, nil
	case "==":
		return c.left.Value == c.right, nil
	case "<=":
		return c.left.Value <= c.right, nil
	case "<":
		return c.left.Value < c.right, nil
	case "!=":
		return c.left.Value != c.right, nil
	default:
		return false, fmt.Errorf("%v unrecognized operator", c.operator)
	}
}
