// package main

// import (
// 	"bytes"
// 	"fmt"
// )

// func handleRequest(){

// 	buf:= new(bytes.Buffer)
// 	buf.WriteString("Processing Request")

// 	fmt.Println(buf.String())
// }
// func main(){

// 	fmt.Println("Sync Pool ")
// 	for i:=0;i<5;i++{
// 		handleRequest()
// 	}
// }

// now with sync pool because above code initialize buffer everytime

package main

import (
	"bytes"
	"fmt"
	"sync"
)



// var bufPool  = sync.Pool{
// 	New: func() any{

// 		return new(bytes.Buffer)
// 	},

// }

var bufPool = sync.Pool{
	New: func() any{
		return new(bytes.Buffer)
	}, 
}

func handleRequest(){
	buf:= bufPool.Get().(*bytes.Buffer)

	buf.Reset()
	buf.WriteString("Processing Request...")
	fmt.Println(buf.String())
	bufPool.Put(buf)
}

func main(){

	n:=4
	for i:=0;i<n;i++{
		handleRequest()
	}


}