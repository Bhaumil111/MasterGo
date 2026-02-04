package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)
var counter int64

func increment ( wg * sync.WaitGroup){
	atomic.AddInt64(&counter, 1)
	defer wg.Done()
}
func main(){

	var wg sync.WaitGroup
	for i:=0 ;i<1000;i++{
		wg.Add(1)

		go increment(&wg)
	}
	wg.Wait()
	fmt.Println("FinalCounter: " ,counter)

}

//====== Compare and Swap

// package main

// import (
// 	"fmt"
// 	"sync/atomic"
// )

// func main(){
// 	var state int64  = 0
// 	fmt.Println("State:" , atomic.LoadInt64(&state))
// 	swapped := atomic.CompareAndSwapInt64(&state, 10 , 1)
// 	fmt.Println("Swapped " , swapped)
// 	fmt.Println("State:" , atomic.LoadInt64(&state))

// }
//============== Swap ( replace and get old value)====


// package main

// import (
// 	"fmt"
// 	"sync/atomic"


// )


// func main(){
// 	var v int64  = 10
// 	old := atomic.SwapInt64(&v , 50)
// 	fmt.Println("Old ",old)
// 	fmt.Println(" New ",atomic.LoadInt64(&v))

	
// }



/// Atomic Bool 
// package main

// import (
// 	"fmt"
// 	"sync/atomic"
// )

// func main() {
// 	var b atomic.Bool

// 	b.Store(true)
// 	fmt.Println("Value:", b.Load())
// }


//== Atomic value=====================


// package main

// import (
// 	"fmt"
// 	"sync/atomic"
// )

// type Config struct {
// 	Name string
// }

// func main() {
// // 	var v atomic.Value

// // 	v.Store(Config{Name: "v1"})
// // 	cfg := v.Load().(Config)

// // 	fmt.Println(cfg.Name)
// var v atomic.Value

// v.Store(Config{Name:"v1"})
// cfg := v.Load().(Config)
// fmt.Println(cfg.Name)
// }
