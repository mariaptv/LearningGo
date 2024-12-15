package main

import (
	"fmt"
	"net/http"
)
const portNumber = ":8080"
//If it starts with uppercase the function is public, else is private 
func Home(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "This is the home page")
}

func About(w http.ResponseWriter, r *http.Request){
	sum := addValues(2,2)
	_,_=fmt.Fprint(w, fmt.Sprintf("This is the home page %d", sum))
}

func addValues(x, y int) int{
	return x + y
}
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	fmt.Println(fmt.Sprintf("Starting app on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
