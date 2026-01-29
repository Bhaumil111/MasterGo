package main

import (
	"fmt"
)

type transformFn func(int) int
type transformFnArray func([]int) []int

func triple(number int) int {
	return number * 3
}

func double(number int) int {
	return number * 2
}

func transformationFunction(numbers []int, transform transformFn) []int {
	doubleNumbers := []int{}
	for _, val := range numbers {
		doubleNumbers = append(doubleNumbers, transform(val))

	}
	return doubleNumbers

}

func getFunction(numbers []int) transformFn {
	if numbers[0] == 1 {
		return double
	} else {

		return triple
	}
}

func printLine(topic string) {
	fmt.Printf("==== %s ======\n", topic)
}

func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}

func sumup(numbers ...int) int {

	sum := 0
	for _, v := range numbers {
		sum += v

	}
	return sum
}

func main() {
	printLine("Learning Functions")
	printLine("Passing Functions as value ")

	nums := []int{1, 2, 3, 4, 5}
	doubleNums := transformationFunction(nums, double)

	tripleNums := transformationFunction(nums, triple)

	fmt.Println(doubleNums)
	fmt.Println(tripleNums)
	printLine("Return type is Function")

	moreNumbers := []int{3, 2, 3, 3}
	func1 := getFunction(moreNumbers)
	fmt.Println(func1)
	transformArray := transformationFunction(moreNumbers, func1)
	// println(transformArray)
	fmt.Println(transformArray)

	printLine("Anonymous Function")

	numbers := []int{3, 2, 3, 4, 5}
	AnonymousArray := transformationFunction(numbers, func(number int) int {
		return number * 2
	})

	fmt.Println(AnonymousArray)

	printLine("Closure Property ")

	doubleClos := createTransformer(2)

	doubleByClos := transformationFunction(numbers, doubleClos)
	fmt.Println(doubleByClos)

	printLine("Varadic Functions")

	// var_nums := []int{3,2,3,3,3}

	sum := sumup(1,2,191,29,19)
	fmt.Println(sum)

}
