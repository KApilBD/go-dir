package main

import (
	"fmt"
	"html/template"
	"net/http"

	"fm.com/go/museum/api"
	"fm.com/go/museum/data"
)

func handleHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the Go paramz!!!"))
}

func handleTemplete(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("templates/index.tmpl")
	if err != nil {
		w.Write([]byte("Opps there is some errorrr!"))
		w.WriteHeader(http.StatusInternalServerError)
	}
	html.Execute(w, data.GetAll()[1])
}

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/hello", handleHello)
	server.HandleFunc("/tmpl", handleTemplete)
	server.HandleFunc("/api/exhibitions", api.Get)
	server.HandleFunc("/api/exhibitions/new", api.Post)

	fs := http.FileServer(http.Dir("./public"))

	server.Handle("/", fs)

	err := http.ListenAndServe(":3333", server)
	if err == nil {
		fmt.Println("Error while opening the server!!!")
	}
}
