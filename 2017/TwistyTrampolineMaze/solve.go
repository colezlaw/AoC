package main

import (
	"io/ioutil"
	"strings"
	"fmt"
	"os"
	"strconv"
)

func input() ([]string, error) {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(content), "\n")
	return lines, nil
}

func main(){
	instructions, err := input()

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	var offsets []int

	for _,v := range instructions{
		value, _ := strconv.Atoi(v)
		offsets = append(offsets, value)
	}

	normal := make([]int, len(offsets))
	copy(normal, offsets)
	steps := FollowInstructions(normal, false)
	fmt.Printf("Steps to reach the exit: %v\n", steps)

	steps = FollowInstructions(offsets, true)
	fmt.Printf("Strange steps to reach the exit: %v\n", steps)


}

func FollowInstructions(offsets []int, stranger bool) int {
	var current, steps int
	for current < len(offsets) {
		jump := offsets[current]
		if stranger && jump >= 3 {
			offsets[current] -= 1
		}else{
			offsets[current] += 1
		}

		current += jump
		steps += 1
	}
	return steps
}
