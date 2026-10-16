// package main

// import (
// 	"fmt"
// 	"net/http"
// )
// //url shortener

// func sayHello(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "hello")
// }

// // here r *http.Request has

// // method
// // URL
// // headers
// // query parameters
// // body
// // cookies

// // ex:
// // r.Method
// // r.URL
// // r.Header
// // r.Body

// func main(){
// 	http.HandleFunc("/", sayHello)
// 	fmt.Println("Server running on http://localhost:8080")
// 	http.ListenAndServe(":8080", nil)
// }

package main

import (
	"fmt"
	"net/http"
	"goshort/internal"
)
func main() {
	internal.InitRoutes()
	fmt.Println("Server running on 8080")
	http.ListenAndServe(":8080", nil)
}
