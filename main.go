/* package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, " Hello World")
}
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, " I am saiful. I am learning backend by golang")
}
func main() {
	mux := http.NewServeMux()


    //mux.HandleFunc("/hey", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintln(w, "Welcome te the server")
	//}
    //// we can also create route like this instend of separate route method.



	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/about", aboutHandler)

	fmt.Println("Server is running on port: 3000")

	err := http.ListenAndServe(":3000", mux)

	if err != nil {
		fmt.Print("Error starting the server", err)

	}
}
