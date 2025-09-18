package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/rasibn/hello.rasib.me/views"
	"github.com/rasibn/hello.rasib.me/views/blog"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./public/assets"))))
	http.HandleFunc("/favicon.ico", icon)
	http.HandleFunc("/", root)
	http.HandleFunc("/weeb", weeb)
	http.HandleFunc("/blog", blogList)
	http.HandleFunc("/blog/", blogPost)

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

func blogList(w http.ResponseWriter, r *http.Request) {
	template := blog.BlogList(Posts)
	_ = template.Render(context.Background(), w)
}

func blogPost(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/blog/")
	id := strings.TrimSuffix(path, "/")

	for _, post := range Posts {
		if post.ID == id {
			template := blog.BlogWrapper(post)
			_ = template.Render(context.Background(), w)
			return
		}
	}

	http.NotFound(w, r)
}
