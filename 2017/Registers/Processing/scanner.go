package Processing

import (
	"os"
	"bufio"
	"strings"
	"strconv"
)

type Scanner struct {
	Instructions Instructions
	Memory Memory
}

func (s *Scanner) Scan() {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scan(file, s)
}

func scan(file *os.File, s *Scanner) {
	scanner := bufio.NewScanner(file)
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
	instruction := Instruction{register, command, value, &condition}

	s.Instructions.Series = append(s.Instructions.Series, &instruction)
}
