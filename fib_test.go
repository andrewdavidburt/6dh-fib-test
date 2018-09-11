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

//func TestFib(t *testing.T) {
//	var result = Fib()
//	var comparison = Fib()
//	if result == nil {
//		t.Error("Expected 0, got ", result)
//	}
//}

func TestIterateFib(t *testing.T) {
	var digits int = 5	
	var comparison = []int{1, 1, 2, 3, 5}
	var result = IterateFib(digits)
	if !(Equal(result, comparison)) {
		t.Error("Expected", comparison, ", received", result)
	}
}

//func TestHandleCall(t *testing.T) {
//	router := mux.NewRouter().StrictSlash(true)
//        router.HandleFunc("/{digits}", HandleCall)
//	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
//	w := httptest.NewRecorder()
//}
