package main

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/f-bh/ghq/handlers"
	"github.com/f-bh/ghq/model"
	"github.com/f-bh/ghq/templates"
)

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./static"))

	state := model.ServerState{}

	scaffold := templates.Scaffold(templates.Lobby())

	mux.Handle("/", templ.Handler(scaffold))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.HandleFunc("/newgame/", handlers.NewSession(&state))

	http.ListenAndServe(":3000", mux)
}
