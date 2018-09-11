package main

import (
	"testing"
//	"net/http/httptest"
//	"net/http"
//	"fmt"
//	"io"
//	"io/ioutil"
//	"log"
)

//slice comparison function, as slices can't be compared with standard operators
func Equal(a, b []int) bool {
    if len(a) != len(b) {
        return false
    }
    for i, v := range a {
        if v != b[i] {
            return false
        }
    }
    return true
}

//test Fib - After failing to figure out how to compare the value of the enclosed
// function output, I then tried to resort to simply comparing the output of two
// identically called functions in order to at least test consistent output, but
// am still getting errors.

//func TestFib(t *testing.T) {
//	var result = Fib()
//	var comparison = Fib()
//	if result == nil {
//		t.Error("Expected 0, got ", result)
//	}
//}

//test IterateFib with custom slice comparison function. This test works.

func TestIterateFib(t *testing.T) {
	var digits int = 5	
	var comparison = []int{1, 1, 2, 3, 5}
	var result = IterateFib(digits)
	if !(Equal(result, comparison)) {
		t.Error("Expected", comparison, ", received", result)
	}
}

//test HandleCall -  This requires spinning up a test router and then sending it
// a sample API call and reading the results. I've been studying several approaches
// that require 3rd party libraries, and haven't finished this in time.

//func TestHandleCall(t *testing.T) {
//	router := mux.NewRouter().StrictSlash(true)
//        router.HandleFunc("/api/fibonacci/{digits}", HandleCall)
//	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
//	w := httptest.NewRecorder()
//}
