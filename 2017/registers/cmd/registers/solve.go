package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/bubba-h57/AoC/2017/registers"
)

var fn *string

func init() {
	fn = flag.String("fn", "./input.txt", "filename to open")
	flag.Parse()
}

func main() {
	f, err := os.Open(*fn)
	if err != nil {
		log.Fatalf("unable to open %v", err)
	}
	defer f.Close()

	scanner := registers.NewScanner(f)

	scanner.Scan()
	scanner.Run()
	fmt.Println(scanner.FindLargestValue())
	fmt.Println(scanner.HighestValue)
}
