// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type Person struct {
// 	Name string `json:"name"`
// 	Age  int    `json:"age"`
// }

// func main() {
// 	p := Person{
// 		Name: "Aditya",
// 		Age:  22,
// 	}

// 	// Convert struct → JSON bytes
// 	data, err := json.Marshal(p)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(string(data))
// }

////// unmarhsl


 package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	input := []byte(`{"name":"Aditya","age":22}`)

	var p Person

	// JSON → struct
	err := json.Unmarshal(input, &p)
	if err != nil {
		panic(err)
	}

	fmt.Println(p.Name)
	fmt.Println(p.Age)
}
