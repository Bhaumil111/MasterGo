package main

import (
	"fmt"
)

// Stage 1: Generate numbers
func generate(nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		for _, n := range nums {
			fmt.Println("Generate:", n)
			out <- n
		}
		close(out)
	}()

	return out
}

// Stage 2: Square numbers
func square(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for n := range in {
			result := n * n
			fmt.Println("Square:", n, "->", result)
			out <- result
		}
		close(out)
	}()

	return out
}

// Stage 3: Print results
func printer(in <-chan int) {
	for n := range in {
		fmt.Println("Print:", n)
	}
}

func main() {
	// Build the pipeline
	stage1 := generate(1, 2, 3, 4)
	stage2 := square(stage1)
	printer(stage2)
}
