// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// func apiHandler(w http.ResponseWriter , r * http.Request ){
// 	fmt.Fprintf(w,"Hello from bhaumil ")

// }
// func main(){
// 	http.HandleFunc("/api" , apiHandler)
// 	fmt.Println("Starting server at port 8000")
// 	http.ListenAndServe(":8000" , nil)
// }

// /////////
// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"time"
// )

// func  headerMiddleware(next http.Handler) http.Handler{
// 	return  http.HandlerFunc(func(w http.ResponseWriter , r * http.Request){
// 		w.Header().Set("X-Custom-Header", "Pokemon")
// 		next.ServeHTTP(w,r)
// 	})
// }

// func loggingMiddleware(next http.Handler) http.Handler {
//     return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//         start := time.Now()
//       next.ServeHTTP(w, r)
//         log.Printf("%s %s %s", r.Method, r.RequestURI, time.Since(start))
//     })
// }

// func homeLander(w http.ResponseWriter, r * http.Request){
// 	fmt.Fprintln(w,"Welcome to the Home Page!")

// }

// func main(){
// 	mux:= http.NewServeMux()

// 	mux.Handle("/" , loggingMiddleware(headerMiddleware(http.HandlerFunc(homeLander))))

// 	mux.Handle("/api"  , loggingMiddleware(headerMiddleware(http.HandlerFunc(homeLander))))

// 	log.Println("Starting server on port 8080")

// 	if err := http.ListenAndServe(":8000",mux ); err != nil{
// 		log.Fatal("Server failed!" , err)
// 	}

// }

/////////

package main

import (
	"fmt"
	"net/http"
	"strings"
)


// https://api.example.com/api/v1/greet?name=Aditya
func greetHandler(w  http.ResponseWriter, r * http.Request){

	query := r.URL.Query()
	fmt.Fprintf(w, "Query %v" , query)

	name := query.Get("name")

	if name== ""{
		name = "Guest"

		fmt.Fprintf(w, "Hello , %s!", name)

	}
}

// https://api.example.com/user/123

func userHandler(w http.ResponseWriter  , r * http.Request){
	pathSegments := strings.Split(r.URL.Path, "/")
	if len(pathSegments) >=3 &&pathSegments[1] =="user"{
		userId := pathSegments[2]
		fmt.Fprintf(w, "User Id: %s", userId)
	}else{
		http.NotFound(w,r)
	}

}





//https://api.example.com/username/123?includeDetails=true


func userDetailsHandler(w http.ResponseWriter, r * http.Request){
	pathSegments :=strings.Split(r.URL.Path , "/")

	query := r.URL.Query()


	includeDetails :=query.Get("includeDetails")

	if len(pathSegments) >=3 && pathSegments[1]=="username"{
		userID :=pathSegments[2]

		response := fmt.Sprintf("User ID: %s" , userID)
		if includeDetails=="true"{
			response +="(Details included)"
		}

		fmt.Fprintln(w, response)
	}else{
		http.NotFound(w,r)
	}
}


func main(){
	http.HandleFunc("/greet" , greetHandler)
	http.HandleFunc("/user/", userHandler)
	http.HandleFunc("/username/" , userDetailsHandler)


	fmt.Println("Listening at port 8080...")
	if err:= http.ListenAndServe(":8080" ,nil); err!=nil{
		fmt.Printf("Failed to listen at port 8080", err)
	}
}