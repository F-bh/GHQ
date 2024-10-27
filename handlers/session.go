package handlers

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/f-bh/ghq/model"
	"github.com/f-bh/ghq/templates"
)

const BASEURL string = "http://localhost:3000"

func NewSession(server *model.ServerState) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		player := model.Player()
		game := model.NewGamestate()
		server.Games = append(server.Games, &game)

		game.Players[0] = player

		if err := r.ParseForm(); err != nil {
			panic("failed to parse session form")
		}

		player.SetDisplayName(r.Form.Get("display-name"))

		names := make([]string, 0, 2)
		for _, player := range game.Players {
			if player != nil {
				names = append(names, player.GetDisplayName())
			}
		}

		w.Header().Add("HX-Replace-Url", "/lobby/"+game.GetSession())

		templ.Handler(
			templates.OpenLobby(
				names, game.GetSession(), BASEURL)).
			ServeHTTP(w, r)
	}
}

func JoinSession(server *model.ServerState) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var game model.IGameState
		player := model.Player()

		println("join hit")

		if err := r.ParseForm(); err != nil {
			panic("failed to parse session form")
		}

		player.SetDisplayName(r.Form.Get("display-name"))

		for _, gameSession := range server.Games {
			if gameSession.GetSession() == r.PathValue("session") {
				game = gameSession
				break
			}
		}

		templ.Handler(
			templates.Scaffold(
				templates.JoinLobby(
					game, BASEURL))).
			ServeHTTP(w, r)
	}
}

func JoinedSession(server *model.ServerState) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		println("joined hit")
		var game model.IGameState
		player := model.Player()

		if err := r.ParseForm(); err != nil {
			panic("failed to parse session form")
		}

		player.SetDisplayName(r.Form.Get("display-name"))

		for _, gameSession := range server.Games {
			if gameSession.GetSession() == r.PathValue("session") {
				game = gameSession
				break
			}
		}

		names := make([]string, 0, 2)
		for _, player := range game.GetPlayers() {
			if player != nil {
				names = append(names, player.GetDisplayName())
			}
		}

		//w.Header().Add("HX-Replace-Url", "/lobby/"+game.GetSession())

		templ.Handler(
			templates.Scaffold(
				templates.OpenLobby(
					names, game.GetSession(), BASEURL))).
			ServeHTTP(w, r)
	}
}
