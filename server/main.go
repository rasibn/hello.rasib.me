package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Static("./static", "frontend/static")
	r.StaticFile("favicon.ico", "frontend/static/favicon_io/favicon.ico")
	r.LoadHTMLGlob("frontend/templates/*.html")

	r.GET("/", root)
	r.POST("/api/posts/:id", getPost)

	r.Run("localhost:8089")
}
