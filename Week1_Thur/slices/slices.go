package main

import "fmt"

func main() {
	//              0   1  2  3  4

	// Real world use case to avoid memory licks
	// 	func readHeader(data []byte) []byte {
	//     return data[:16]
	// }
	// func readHeader(data []byte) []byte {
	//     h := make([]byte, 16)
	//     copy(h, data[:16])
	//     return h
	// }

	// arr := [5]int{10, 20, 30, 40, 50}
	// s := arr[1:4]
	// fmt.Println(s)
	// fmt.Println("I am learning about slices")

	// s := [] int{10,20,30}
	// fmt.Println(s, len(s) , cap(s))

	// s:= []int {1,2,3,4}
	// a:=s[:2]

	// a[0] =99
	// fmt.Println("A", a)
	// fmt.Println("S", s)

	// s := make([]int, 0, 5)
	// a := append(s, 1)
	// b := append(a, 2)

	// fmt.Println("s:", a , len(s), cap(s))
	// fmt.Println("a:", a , len(a), cap(a))
	// fmt.Println("b:", b, len(b), cap(b))

	s := []int{1, 2, 3}
	a := s[:2] // 1 , 2
	// b := append(a, 100,200)
	b := a[:2:4]
	b = append(b, 100, 400)
	fmt.Println("s:", s, len(s), cap(s))
	fmt.Println("a:", a, len(a), cap(a))
	fmt.Println("b:", b, len(b), cap(b))

}
