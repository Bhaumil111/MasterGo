// package main

// import (
// 	"fmt"
// 	"sync"
// )

// var counter = 0

// func increment(wg *sync.WaitGroup) {
// 	counter++
// 	wg.Done()
// }

// func main() {
// 	var wg sync.WaitGroup
// 	expectedCounter := 1000

// 	for i := 0; i < expectedCounter; i++ {
// 		wg.Add(1)
// 		go increment(&wg)
// 	}

// 	wg.Wait()
// 	fmt.Println("Expected Counter:", expectedCounter)
// 	fmt.Println("Actual Counter:", counter)
// 	// Check for race condition
// 	if expectedCounter != counter {
// 		fmt.Println("Race condition detected!")
// 	} else {
// 		fmt.Println("No race condition detected.")
// 	}
// }

package main

import (
	"fmt"
	"sync"
)
var counter =0
var mu sync.Mutex
func increment(wg * sync.WaitGroup){


	defer wg.Done()
	// mu.Lock()
	mu.Lock()
	counter++
	defer mu.Unlock()
}
func main(){

	var wg sync.WaitGroup

	expcounter :=1000
	for i:=0 ;i<expcounter;i++{
		wg.Add(1)

		go increment(&wg)


	}
	wg.Wait()

	fmt.Println("Exp counter", expcounter)
	fmt.Println("Actual Counter: " , counter)
	if(expcounter!=counter){
		fmt.Println("Race condition detected")
	}else{
		fmt.Println("No Race condition detected")
	}

}