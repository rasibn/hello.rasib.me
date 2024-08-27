SESSION := WEBSITE

install:
	go install github.com/air-verse/air@latest
	npm install # install tailwindcss
	cp  node_modules/htmx.org/dist/htmx.min.js frontend/static/lib/htmx.min.js # install htmx
	cd server
	go mod tidy # install all the go go modules
	cd ..

go-release:
	go build  -ldflags="-s -w" -o build/Release ./server

tailwind:
	./node_modules/.bin/tailwindcss -i frontend/source.css -o frontend/static/tailwind.css --watch

build-tailwind:
	./node_modules/.bin/tailwindcss -i frontend/source.css -o frontend/static/tailwind.css
	
templ:
	templ generate -watch -proxy=http://localhost:8089

install-tailwind:
	npm install tailwindcss@next @tailwindcss/cli@next

tmux-setup:
	tmux new-session -d -s $(SESSION) -n code 'nvim'
	tmux new-window -t $(SESSION):2 -n zsh
	tmux new-window -t $(SESSION):3 -n air 'air'
	tmux new-window -t $(SESSION):4 -n tailwind 'make tailwind'
	tmux new-window -t $(SESSION):5 -n templ 'make templ'
	tmux select-window -t $(SESSION):1
	tmux attach -t $(SESSION)
