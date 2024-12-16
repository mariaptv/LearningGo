package mainPractice

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// If it starts with uppercase the function is public, else is private
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is the home page")
}

func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	_, _ = fmt.Fprint(w, fmt.Sprintf("This is the home page %d", sum))
}

func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(100.0, 10.0)
	if err != nil {
		fmt.Fprintf(w, "Cannot divide by 0")
		return
	}

	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", 100.0, 10.0, f))
}

func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("Cannot divide by zero")
		return 0, err
	}
	result := x / y
	return result, nil
}

func addValues(x, y int) int {
	return x + y
}
func mainPractice() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)
	fmt.Println(fmt.Sprintf("Starting app on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
