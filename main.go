package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, nil)
}

func main() {
	port := ":8000"
	fmt.Println("Web server starting")
	http.HandleFunc("/", indexHandler)
	fmt.Println("Now serving requests at port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
