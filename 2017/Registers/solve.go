package main

import (
	"fmt"
	"./Processing"
)

func main()  {
	var scanner Processing.Scanner
	scanner.Scan()
	scanner.Instructions.Run()
	fmt.Println(scanner.Memory.FindLargestValue())
	fmt.Println(scanner.Memory.HighestValue)
}
