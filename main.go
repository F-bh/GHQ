package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/f-bh/ghq/handlers"
	"github.com/f-bh/ghq/model"
	"github.com/f-bh/ghq/templates"
)

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./static"))

	state := model.ServerState{
		BaseUrl: "localhost:3000",
	}

	createLobby := templates.Scaffold(templates.Lobby())

	mux.Handle("/", templ.Handler(createLobby))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/newgame/", handlers.NewSession(&state))
	mux.HandleFunc("/lobby/{session}", handlers.JoinSession(&state))
	mux.HandleFunc("/lobby/join/{session}", handlers.JoinSession(&state))
	mux.HandleFunc("/lobby/joined/{session}", handlers.JoinedSession(&state))
	mux.HandleFunc("/events/{session}", handlers.EventHandler(&state))

	fmt.Println("serving on :3000")
	http.ListenAndServe(":3000", mux)
}
