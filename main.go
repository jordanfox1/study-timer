package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var (
	port = ":8080"
)

type Task struct {
	Title    string
	Priority string
}

func main() {
	handler := func(w http.ResponseWriter, req *http.Request) {
		templ := template.Must(template.ParseFiles("index.html"))
		templ.Execute(w, nil)
	}

	handler2 := func(w http.ResponseWriter, r *http.Request) {
		title := r.PostFormValue("title")
		priority := r.PostFormValue("priority")
		templ := template.Must(template.ParseFiles("index.html"))
		templ.ExecuteTemplate(w, "task-list-element", Task{Title: title, Priority: priority})

	}

	http.HandleFunc("/", handler)
	http.HandleFunc("/add-task/", handler2)
	fmt.Printf("server will start on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
