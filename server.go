package main

import (
	"html/template"
    "net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { //crée une page avec n'importe quel URL 
		tmpl, err := template.ParseFiles("./tmpl/index.html")
		tmpl.ExecuteTemplate(w, "index", nil)
		if err != nil {
			error501()
		}
	})
	http.ListenAndServe("localhost:3000", nil)
}

func error404() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { //crée une page avec n'importe quel URL si il y a une erreur 404
		tmpl, err := template.ParseFiles("./error404.html")
		tmpl.ExecuteTemplate(w, "error404", nil)
		if err != nil {
			error501()
		}
	})
	http.ListenAndServe("localhost:3000", nil)
}

func error501() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { //crée une page avec n'importe quel URL si il y a une erreur 501
		tmpl, err := template.ParseFiles("./error501.html")
		tmpl.ExecuteTemplate(w, "error501", nil)
		if err != nil {
			// error501()
		}
	})
	http.ListenAndServe("localhost:3000", nil)
}
func remove(slice []string, s int) []string {
	/*
	fonction permettant de retirer n'importe quel élément a l'index "s" que l'on souhaite
	*/
    return append(slice[:s], slice[s+1:]...)
}
