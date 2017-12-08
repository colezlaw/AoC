package registers

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type Scanner struct {
	Instructions
	Memory

	r io.Reader
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: r}
}

func (s *Scanner) Scan() {
	scanner := bufio.NewScanner(s.r)
	for scanner.Scan() {
		parse(scanner, s)
	}
}

func parse(scanner *bufio.Scanner, s *Scanner) {
	fields := strings.Fields(scanner.Text())

	left := s.Memory.Add(fields[4])
	register := s.Memory.Add(fields[0])
	command := fields[1]
	operator := fields[5]
	value, _ := strconv.Atoi(fields[2])
	right, _ := strconv.Atoi(fields[6])

	condition := Condition{left, operator, right}
	instruction := instruction{register, command, value, &condition}

	s.Instructions.Series = append(s.Instructions.Series, &instruction)
}
