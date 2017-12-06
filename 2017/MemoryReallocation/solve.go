package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func input() ([]int) {
	file, _ := os.Open("./input.txt")
	defer file.Close()
	var banks []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()
		for _, field := range strings.Fields(input) {
			parsed, _ := strconv.Atoi(field)
			banks = append(banks, parsed)
		}
	}
	return banks
}


func main() {
	state := input()
	
	redistributions := findInfiniteLoop(state, false)
	fmt.Printf("Infinite Loop Detected after %v redistribution cycles.\n", redistributions)


	loopCycles := findInfiniteLoop(state, true)
	fmt.Printf("%v cycles in the Infinite Loop.\n", loopCycles)
}


func findInfiniteLoop(state []int, loopCycles bool) int {
	cycles1 := map[string]int{}
	cycles2 := map[string]bool{}

	redistributions := 0
	for {
		stringified := strings.Join(strings.Fields(fmt.Sprint(state)), " ")

		if loopCycles {
			if cycles1[stringified] > 0 {
				return len(cycles1) - (cycles1[stringified] - 1)
			}

			cycles1[stringified] = redistributions + 1
		}else{
			if cycles2[stringified] {
				break
			}

			cycles2[stringified] = true
		}


		idx := 0
		max := math.MinInt32
		for maxIdx, val := range state {
			if val > max {
				max = val
				idx = maxIdx
			}
		}

		val := state[idx]
		state[idx] = 0

		for i := 0; i < val; i++ {
			next := (idx + i + 1) % len(state)

			state[next] = state[next] + 1
		}

		redistributions += 1
	}
	return redistributions
}
