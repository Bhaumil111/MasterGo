// ============================= With Timeout

// package main

// import (
// 	"context"
// 	"fmt"
// 	"time"
// )

// func main() {
//  ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
//  defer cancel()

//  go performTask(ctx)

//  select {
//  case <-ctx.Done():
//   fmt.Println("Task timed out")
//  }
//  time.Sleep(5*time.Second)
//  fmt.Println("hi ")
// }

// func performTask(ctx context.Context) {
// 	// time.Sleep(4*time.Second)
//  select {
//  case <-time.After(5 * time.Second):
//   fmt.Println("Task completed successfully")

//  case<-ctx.Done():
// 	fmt.Println(" Perform task function ", ctx.Err())
// 	// return
//  }
// }

//=======================With value

// package main

// import (
// 	"context"
// 	"fmt"
// )

// func main(){

// 	ctx := context.WithValue(context.Background() , "user_id" , 123)

// 	processRequest(ctx)

// }
// func processRequest(ctx  context.Context){
// 	User_id := ctx.Value("user_id").(int)
// 	fmt.Println("Your user_id is " , User_id)
// }

// =========================With Cancel

// package main

// import (
//  "context"
//  "fmt"
//  "time"
// )

// func main() {
//  ctx, cancel := context.WithCancel(context.Background())

//  go performTask(ctx)

//  time.Sleep(2 * time.Second)
//  cancel()

//  time.Sleep(1 * time.Second)
// }

// func performTask(ctx context.Context) {
//  for {
//   select {
//   case <-ctx.Done():
//    fmt.Println("Task cancelled")
//    return
//   default:
//    // Perform task operation
//    fmt.Println("Performing task...")
//    time.Sleep(500 * time.Millisecond)
//   }
//  }
// }

//=========================With Deadline


package main

import (
 "context"
 "fmt"
 "time"
)

func main() {
 ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))
 defer cancel()

 go performTask(ctx)

 time.Sleep(3 * time.Second)
}

func performTask(ctx context.Context) {
 select {
 case <-ctx.Done():
  fmt.Println("Task completed or deadline exceeded:", ctx.Err())
  return
 }
}