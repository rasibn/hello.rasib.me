package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/rasibn/hello.rasib.me/views"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./public/assets"))))
	http.HandleFunc("/favicon.ico", icon)
	http.HandleFunc("/", root)
	http.HandleFunc("/weeb", weeb)

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
	template := views.Hello()
	_ = template.Render(context.Background(), w)
}

func weeb(w http.ResponseWriter, r *http.Request) {
	template := views.Weeb()
	_ = template.Render(context.Background(), w)
}
