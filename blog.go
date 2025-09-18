package main

import (
	"github.com/rasibn/hello.rasib.me/views/blog"
)

var Posts = []blog.BlogPost{
	{
		ID:       "1",
		Title:    "My First Blog Post",
		Date:     "January 15, 2024",
		Summary:  "Welcome to my first blog post! Sharing thoughts on programming and technology.",
		Draft:    true,
		Template: blog.Blog1Content,
	},
	{
		ID:       "2",
		Title:    "Learning Go: Lessons from Building CLI Tools",
		Date:     "February 2, 2024",
		Summary:  "Why Go works great as a scripting language and lessons from building daily-use CLI tools.",
		Draft:    true,
		Template: blog.Blog2Content,
	},
	{
		ID:       "3",
		Title:    "Building a Reading App: Why I Made Hatsulingo",
		Date:     "March 1, 2024",
		Summary:  "The story behind building hatsulingo.com - a web-based reading app focused on simplicity.",
		Draft:    true,
		Template: blog.Blog3Content,
	},
}
