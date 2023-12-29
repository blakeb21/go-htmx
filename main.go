package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Server Running")

	home := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("home.html"))
		tmpl.Execute(w, nil)
	}

	coaching := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("coaching.html"))
		tmpl.Execute(w, nil)
	}

	http.HandleFunc("/", home)
	http.HandleFunc("/coaching", coaching)

	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./images"))))
	log.Fatal(http.ListenAndServe(":8000", nil))
}
