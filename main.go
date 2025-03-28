package main

import (
	"flag"
	"fmt"
)

// ./euromillions --lines 16
func main() {
	numLines := flag.Int("lines", 3, "Number of Euromillions lines to generate")

	fmt.Println("Generating", *numLines, "lines")
	flag.Parse()

	lines := GenerateLines(*numLines)

	for _, line := range lines {
		fmt.Println(line)
	}
}
 