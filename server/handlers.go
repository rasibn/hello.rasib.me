package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func root(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Title": "My Journey in go-lang(d)",
	})
}

func getPost(c *gin.Context) {
	id := c.Param("id")
	fmt.Printf("Recieved id %s", id)

	c.HTML(http.StatusOK, "index.html", gin.H{})
}
