package main

import "fmt"

func seperator() {
	fmt.Println("============================")
}

func main() {
	items := []int{1, 2, 3, 4, 5, 6}
	batches := distribute(items)
	// Expected: [[1,2,3], [4,5,6]]
	// Got: [[1,2,3], [4,5,6]] looks right...

	batches[0][0] = 999
	fmt.Println(batches) // [[999,2,3], [999,5,6]] - HUH?!

}

// classic loop
func distribute(items []int) [][]int {
	result := [][]int{}
	batch := []int{}

	for _, item := range items {
		batch = append(batch, item)
		if len(batch) == 3 {
			result = append(result, batch)
			batch = []int{} // WRONG!
		}
	}
	return result
}
