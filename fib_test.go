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

func TestFib(t *testing.T) {
	var result = Fib()
	var comparison = Fib()
	if result == nil {
		t.Error("Expected 0, got ", result)
	}
}

//func TestIterateFib(t *testing.T) {
//	digits := 1
////increase digits
//	result := IterateFib(digits)
//	if result != 0 {
//		t.Error("Expected 0, got ", result)
//	}
//}

//func TestHandleCall(t *testing.T) {
//	router := mux.NewRouter().StrictSlash(true)
//        router.HandleFunc("/{digits}", HandleCall)
//	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
//	w := httptest.NewRecorder()
//}
