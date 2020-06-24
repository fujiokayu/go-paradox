package main

import (
	"fmt"
)

func writeNaturalNumber(dist float64, naturalNumber uint64) {
	var reached = dist / 2
	fmt.Printf("Gopher 「%d ! 」\n", naturalNumber)
	naturalNumber++

	writeNaturalNumber(reached, naturalNumber)
}

func main() {
	var dist = 100.0             // 任意の距離
	var naturalNumber uint64 = 0 // 自然数は 0 から数えるらしい

	writeNaturalNumber(dist, naturalNumber)
}
