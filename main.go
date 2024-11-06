package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {

	type data struct {
		Title string
	}

	d := data{"rylimitless"}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method)
		t, err := template.ParseFiles("templates/index.html")
		if err != nil {
			fmt.Fprintf(w, "%v", "Error Parsing data")
		} else {
			t.Execute(w, d)
		}
	})

	http.HandleFunc("/form", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method)
		err := r.ParseForm()
		if err == nil {
			name := r.FormValue("name")
			message := r.FormValue("message")
			email := r.FormValue("email")
			fmt.Println(name)
			fmt.Println(message)
			fmt.Println(email)
		}

	})

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("This is a test")

	})

	http.ListenAndServe(":8089", nil)
}
