// package main

// import (
// 	"encoding/json"
// 	"os"
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

// 	json.NewEncoder(os.Stdout).Encode(p)
// }

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"os"
// )


// type Person struct{
// 	Name string `json:"name"`
// 	Age int `json:"age"`
// }

// func main(){
// 	p:=Person{
// 		Name: "Aditya",
// 		Age:22,
// 	}

// 	json.NewEncoder(os.Stdout).Encode(p)
// }

package main

import (
	"encoding/json"
	"strings"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	input := `{"name":"Aditya","age":22}`

	var p Person

	json.NewDecoder(strings.NewReader(input)).Decode(&p)

	fmt.Println(p.Name, p.Age)
}
