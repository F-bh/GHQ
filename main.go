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

	createLobby := templates.Scaffold(templates.CreateLobby())

	mux.Handle("/", templ.Handler(createLobby))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.HandleFunc("POST /lobby/create", handlers.NewSession(&state))
	mux.HandleFunc("GET /lobby/join/{session}", handlers.JoinSession(&state))
	mux.HandleFunc("/lobby/{session}", handlers.JoinedSession(&state))
	mux.HandleFunc("/events/{session}", handlers.EventHandler(&state))
	mux.HandleFunc("/test", sseHandler)

	fmt.Println("serving on :3000")
	http.ListenAndServe(":3000", mux)
}

func sseHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Client connected")
	w.Header().Set("Access-Control-Allow-Origin", "*") // must have since it enable cors
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	flusher, ok := w.(http.Flusher)
	if !ok {
		fmt.Println("Could not init http.Flusher")
	}

	for {
		select {
		case <-r.Context().Done():
			fmt.Println("Connection closed")
			return
		default:
			fmt.Println("case message... sending message")
			fmt.Fprintf(w, "data: Ping\n\n")
			flusher.Flush()
		}
	}
}
