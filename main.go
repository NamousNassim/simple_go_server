package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "NOT FOUND ", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method error", http.StatusNotAcceptable)
		return
	}

	fmt.Fprintf(w, "hello world")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, " ParseForm error %s", err)
		return
	}
	fmt.Fprint(w, "Post sucess")
	name := r.FormValue("name")
	adress := r.FormValue("adress")

	fmt.Fprintf(w, "name = %s\n", name)
	fmt.Fprintf(w, "adress = %s\n", adress)
}

func main() {
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
