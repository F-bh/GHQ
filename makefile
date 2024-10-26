run:
	templ generate templates/
	npx tailwindcss build -o static/css/tailwind.css
	go run main.go