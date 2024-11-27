package main

import (
	"fmt"
	"io"
	"net/http"
)

func methodHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Printf("\nGot a Get request")
	case "POST":
		fmt.Printf("\nGot a POST request")
	default:
		fmt.Printf("Default case")
	}
}

var num = 0

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("\n%d HOMEPAGE!!!!!!", num)
	methodHandler(w, r)
	num++
	io.WriteString(w, "hi")
	// log.Logger("Homepage accessed")
}

func main() {
	fmt.Println("master server!")
	http.HandleFunc("/", homepage)
	err := http.ListenAndServe(":3333", nil)
	print(err)
}
