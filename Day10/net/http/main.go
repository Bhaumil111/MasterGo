//==========================Simple http code
// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// func apiHandler ( w http.ResponseWriter , r * http.Request){
// 	w.Write([]byte("Hello world"))
// }
// func main(){

// 	http.HandleFunc("/api"  , apiHandler)

// 	fmt.Println("Starting server at port 8000")

// 	http.ListenAndServe(":8000" , nil)

// }

//==========================Simple http code using mux

// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// type HomeLander struct{

// }

// func (h HomeLander) ServeHTTP(w http.ResponseWriter , r *http.Request){

// 	fmt.Fprintf(w,"Welcome to the server")
// }
// func main(){

// 	mux:= http.NewServeMux()

// 	mux.Handle("/" , HomeLander{})
// 	mux.HandleFunc("/hello", func( w http.ResponseWriter, r * http.Request){
// 		fmt.Fprintf(w,"Hello world! You are accesing %v  and using user agent %v\n" ,r.URL.Path , r.Header.Get("User-Agent") )

// 	})

// 	mux.HandleFunc("/status" , func(w http.ResponseWriter, r *http.Request){

// 		w.WriteHeader(http.StatusOK)
// 		fmt.Fprintln(w,`{status: "ok"}`)
// 	})

// 	fmt.Println("Server starting on port 8000 ...")
// 	http.ListenAndServe(":8000", mux)

// }

// ====================================Routing and middleware
// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"time"
// 	// "time"
// )

// func headerMiddleware(next http.Handler) http.Handler {
//     return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//         // Implement logic 

//         w.Header().Set("X-Custom-Header", "Pokemon")

//         // End of middleware logic
//         next.ServeHTTP(w, r)
//     })
// }

// func loggingMiddleware(next http.Handler) http.Handler {
//     return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//         start := time.Now()
//         log.Printf("%s %s %s", r.Method, r.RequestURI, time.Since(start))
//         next.ServeHTTP(w, r)
//     })
// }

// func homeHandler(w http.ResponseWriter, r *http.Request) {
//     fmt.Fprintln(w, "Welcome to the Home Page!")
// }

// func aboutHandler(w http.ResponseWriter, r *http.Request) {
//     fmt.Fprintln(w, "This is the about page!")
// }

// func main() {
//     mux := http.NewServeMux()

//     mux.Handle("/", loggingMiddleware(headerMiddleware(http.HandlerFunc(homeHandler))))
//     mux.Handle("/about", loggingMiddleware(headerMiddleware(http.HandlerFunc(aboutHandler))))

//     log.Println("Starting server on port 8080 ...")
//     if err := http.ListenAndServe(":8080", mux); err != nil {
//         // the logic that should be executed in case the listen adn serve returns error
//         log.Fatal("Server Failed!", err)     
//     }
// }





// =========================Query parameters
// package main

// import (
//     "fmt"
//     "net/http"
//     "strings"
// )

// // https://api.example.com/api/v1/greet?name=Aditya
// // https://api.example.com/api/v1/greet

// // Understanding Query Parameters
// func greetHandler(w http.ResponseWriter, r *http.Request) {
//     query := r.URL.Query()
//     name := query.Get("name")
//     if name == "" {
//         name = "Guest" 
//     }
//     fmt.Fprintf(w, "Hello, %s!", name)
// }

// // Extracting Path Variables

// // https://api.example.com/user/123
// // 1 -> User 
// // 2 -> UserID
// func userHandler(w http.ResponseWriter, r *http.Request) {
//     pathSegments := strings.Split(r.URL.Path, "/")
//     if len(pathSegments) >= 3 && pathSegments[1] == "user" {
//         userID := pathSegments[2]
//         fmt.Fprintf(w, "User ID: %s", userID)
//     } else {
//         http.NotFound(w, r)
//     }
// }

// // Combining Both => Handling Query Parameters + Extracting Path Variables
// // https://api.example.com/username/123?includeDetails=true
// func userDetailsHandler(w http.ResponseWriter, r *http.Request) {
//     pathSegments := strings.Split(r.URL.Path, "/")
//     query := r.URL.Query()
//     includeDetails := query.Get("includeDetails")

//     if len(pathSegments) >= 3 && pathSegments[1] == "username" {
//         userID := pathSegments[2]
//         response := fmt.Sprintf("User ID: %s", userID)
//         if includeDetails == "true" {
//             response += " (Details included)"
//         }
//         fmt.Fprintln(w, response)
//     } else {
//         http.NotFound(w, r)
//     }
// }

// func main() {
//     http.HandleFunc("/greet", greetHandler)
//     http.HandleFunc("/user/", userHandler)
//     http.HandleFunc("/username/", userDetailsHandler)

//     fmt.Println("Listening at port 8080 ...")
//     // We are using default mux for demonstration purposes
//     if err := http.ListenAndServe(":8080", nil); err != nil {
//         fmt.Println("Failed to listen at port 8080", err)
//     }
// }

// https://api.example.com/api/v1/notion?sessionToken="823492384"&referralToken="user123" -> User123 was the one who refered -> referralToken -> +50 Points
// https://api.example.com/api/v1/notion?sessionToken="823492384" -> User came here and signed up by his own.

// https://api.example.com/api/v1/status/order/1234






// ================================================ JSON API 





// package main

// import (
//     "encoding/json"
//     "fmt"
//     "log"
//     "net/http"
//     "sync"
// )

// type User struct {
//     ID      int     `json:"id"`
//     Name    string  `json:"name"`
//     Email   string  `json:"email"`
// }

// var (
//     users = make(map[int]User)
//     idSeq = 1
//     mutex = &sync.Mutex{}
// )

// func usersHandler(w http.ResponseWriter, r *http.Request) {
//     w.Header().Set("Content-Type", "application/json")

//     switch r.Method {
//         case "GET":
//             mutex.Lock()
//             defer mutex.Unlock()

//             usersList := make([]User, 0, len(users))
//             for _, user := range users {
//                 usersList = append(usersList, user)
//             }

//             json.NewEncoder(w).Encode(usersList)

//         case "POST":
//             var user User
//             if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
//                 http.Error(w, "Invalid JSON", http.StatusBadRequest)
//                 return
//             }

//             mutex.Lock()
//             user.ID = idSeq
//             idSeq++
//             users[user.ID] = user 
//             mutex.Unlock()

//             w.WriteHeader(http.StatusCreated)
//             json.NewEncoder(w).Encode(user)

//         default:
//             http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
//     }
// }

// func userHandler(w http.ResponseWriter, r *http.Request) {
//     w.Header().Set("Content-Type", "application/json")

//     var id int 
//     _, err := fmt.Sscanf(r.URL.Path, "/users/%d", &id)
//     if err != nil {
//         http.Error(w, "Invalid User ID", http.StatusBadRequest)
//         return
//     }

//     mutex.Lock()
//     defer mutex.Unlock()

//     user, exists := users[id]
//     if !exists {
//         http.Error(w, "User Not Found", http.StatusNotFound)
//         return
//     }

//     switch r.Method {
//         case "GET":
//             json.NewEncoder(w).Encode(user)

//         case "PUT":
//             var updatedUser User
//             if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
//                 http.Error(w, "Invalid JSON", http.StatusBadRequest)
//                 return 
//             }

//             updatedUser.ID = id 
//             users[id] = updatedUser
//             json.NewEncoder(w).Encode(updatedUser)

//         case "DELETE":
//             delete(users, id)
//             w.WriteHeader(http.StatusNoContent)

//         default: 
//             http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
//     }
// }

// func main() {
//     http.HandleFunc("/users", usersHandler)
//     http.HandleFunc("/users/", userHandler)

//     fmt.Println("Server running on port 8080")
//     log.Fatal(http.ListenAndServe(":8080", nil))
// }




// ==== Working Json Practice



package main
import "fmt"





func main(){
	fmt.Println("Now working with net/http")


}