package main

import (
	"fmt"
	"net/http"
)

var todos []string

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Todo List</h1>")
	fmt.Fprintf(w, "<form action=\"/add\" method=\"POST\">")
	fmt.Fprintf(w, "<input type=\"text\" name=\"todo\">")
	fmt.Fprintf(w, "<input type=\"submit\" value=\"Add\">")
	fmt.Fprintf(w, "</form>")
	fmt.Fprintf(w, "<ul>")
	for _, todo := range todos {
		fmt.Fprintf(w, "<li>%s</li>", todo)
	}
	fmt.Fprintf(w, "</ul>")
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	todos = append(todos, r.FormValue("todo"))
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/add", addHandler)
	http.ListenAndServe(":8080", nil)
}
