package main

import (
	"fmt"
	"strconv"
	//"encoding/json"
	//"github.com/manyminds/api2go"
	//"github.com/julienschmidt/httprouter"
	"net/http"
	"github.com/gorilla/mux"
	"log"
	// "testing"
)

// Fib returns a function that returns fibonacci #s
func Fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

// IterateFib calls Fib as many times as is specified in API call
func IterateFib(n int) []int {
	var f = Fib()
	var output []int

        for i := 0; i < n; i++ {
                output = append(output, f())
        }
	
	return output
}

// HandleCall handles the API call, using URL variable 
// to call IterateFib and return the sequence
func HandleCall(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	digits, _ := strconv.Atoi(vars["digits"])
	output := IterateFib(digits)
//	outJ, _ := json.Marshal(IterateFib(digits))

//	fmt.Fprintln(w, "json fib show:", outJ)
	fmt.Fprintln(w, "{\"fibonacci\":", output, "}")
}

func main() {

//	outJ, _ := json.Marshal(output)
//	fmt.Println(string(outJ))

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/{digits}", HandleCall)
	log.Fatal(http.ListenAndServe(":8080", router))
}
