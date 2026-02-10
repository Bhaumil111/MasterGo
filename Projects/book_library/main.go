package main

import (
	"fmt"
	"net/http"
)


func HomeLibrary(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Welcome from home route")
}


// main go -- mux 




func main() {


	mux := http.NewServeMux()

	mux.HandleFunc("GET /", HomeLibrary)



	


	fmt.Println("Hello world")

}
