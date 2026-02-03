// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func recurse(n int, wg *sync.WaitGroup) {
// 	fmt.Println("Enter level", n)

// 	if n == 0 {
// 		fmt.Println("Level 0 doing work")
// 		time.Sleep(1 * time.Second)
// 		wg.Done()
// 		wg.Done()
// 		wg.Done()
// 		wg.Done()
// 		wg.Done()
// 		return
// 	}

// 	wg.Add(1)

// 	go recurse(n-1, wg)

// 	fmt.Println("Parent", n, "waiting")
// 	wg.Wait()
// 	fmt.Println("Parent", n, "response")
// 	// wg.Done()
// }

// func test(wg *sync.WaitGroup) {
// 	wg.Add(1)
// 	go func() {
// 		time.Sleep(time.Microsecond * 550)
// 		wg.Done()
// 		fmt.Println("In test function")
// 	}()
// 	wg.Wait()
// }

// func main() {
// 	var wg sync.WaitGroup

// 	go test(&wg)
// 	wg.Add(1)
// 	// go recurse(3, &wg)
// 	go func() {
// 		time.Sleep(time.Microsecond * 500)
// 		wg.Done()
// 	}()

// 	wg.Wait()
// 	fmt.Println("Main done")
// }

// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func worker ( id int){
// 	fmt.Printf("Worker %d starting\n " , id)

// 	time.Sleep(time.Second)
// 	fmt.Printf("Worker %d done \n" , id)
// }

// func main(){

// 	var wg sync.WaitGroup

// 	for i:=1;i<3;i++{
// 		wg.Go(func(){
// 			worker(i)
// 		})

// 	}
// 	fmt.Print("Main is waiting")
// 	wg.Wait()
// 	fmt.Printf("Main is finished")

// }
// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main(){

// 	var wg sync.WaitGroup

// 	wg.Go(func(){
// 		fmt.Println("Hi i am working on go ")
// 	})

// 	wg.Go(func(){

// 		fmt.Println("Hi I am on working on go 2")
// 	})

// 	wg.Wait()

// }
//=========================================================
// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func worker(id int){
// 	fmt.Println("Hello from " ,id)
// 	time.Sleep(time.Second)
// 	fmt.Println("Done from " ,id)

// }

// func main(){

// 	var wg sync.WaitGroup

// 	i := 0
// 	for i= 0;i<5;i++{

// 		// j:=i

// 		wg.Go(func(){
// 			worker(i)
// 		})
// 	}

// 	wg.Wait()
// 	fmt.Println("Main function done`")

// }

//=========================================
// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )
// func main() {
// 	var wg  sync.WaitGroup

// 	wg.Go(func() {
// 		time.Sleep(1 * time.Second)
// 		fmt.Println("Worker done")
// 	})

// 	go func (){

// 		wg.Wait()
// 	}()

// 	fmt.Println("Main waiting")
// 	wg.Wait()
// }
//==================== Deadlock===========================

// package main

// import (
// 	"fmt"
// 	"sync"
// 	// "time"
// )

// func work(wg * sync.WaitGroup){
// 	fmt.Println("Worker started")
// 	fmt.Println("Work done")
// 	// wg.Done()
// 	// wg.Wait()
// }
// func main() {
// 	var wg sync.WaitGroup

// 	wg.Go(func(){

// 		work(&wg)

// 	})
// 	fmt.Println(
// 		"Main function waiting")
// 	wg.Wait()

// 	fmt.Println("Main function done")

// }

// ========================== Recursive Flow

package main

import (
	"fmt"
	"sync"
)

func recurse(n int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Enter level", n)

	if n == 0 {
		fmt.Println("Base work at level 0")
		return
	}

	// ❌ forgot wg.Add(1)
	go recurse(n-1, wg)

	fmt.Println("Exit level", n)
}


func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go recurse(3, &wg)

	wg.Wait()
	fmt.Println("Main done")
}

