package blog

var Posts = []BlogPost{
	{
		ID:       "1",
		Title:    "Blog Goals",
		Date:     "September 17, 2025",
		Summary:  "Welcome to my Blog where I can plan to write",
		Draft:    false,
		Template: Blog1Content,
	},
	{
		ID:       "2",
		Title:    "Learning Go: Lessons from Building CLI Tools",
		Date:     "February 2, 2024",
		Summary:  "Why Go works great as a scripting language and lessons from building daily-use CLI tools.",
		Draft:    true,
		Template: Blog2Content,
	},
	{
		ID:       "3",
		Title:    "Building a Reading App: Why I Made Hatsulingo",
		Date:     "March 1, 2024",
		Summary:  "The story behind building hatsulingo.com - a web-based reading app focused on simplicity.",
		Draft:    true,
		Template: Blog3Content,
	},
}
