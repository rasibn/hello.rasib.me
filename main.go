package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./public/assets"))))
	http.HandleFunc("/favicon.ico", icon)
	http.HandleFunc("/", root)

	url := "0.0.0.0:8089"
	fmt.Println("Server is running on ", url)
	if err := http.ListenAndServe(url, nil); err != nil {
		log.Fatal(err)
	}
}

func icon(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, filepath.Join("public", "assets", "favicon_io", "favicon.ico"))
}

func root(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseGlob("./public/templates/*.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	type Data struct {
		Title string
	}

	data := Data{
		Title: "hello world",
	}

	if err = tmpl.ExecuteTemplate(w, "index.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
