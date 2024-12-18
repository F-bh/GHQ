package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/f-bh/ghq/model"
	"github.com/f-bh/ghq/templates"
)

func NewSession(server *model.ServerState) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		player := model.Player()
		game := model.NewGamestate()
		server.Games = append(server.Games, game)

		game.Players[0] = player

		if err := r.ParseForm(); err != nil {
			log.Println("failed to parse session form")
		}

		player.SetDisplayName(r.Form.Get("display-name"))

		joinUrl := templ.URL(fmt.Sprintf("http://%v/lobby/join/%v", server.BaseUrl, game.GetSessionId()))

		w.Header().Add("HX-Replace-Url", "/lobby/"+game.GetSessionId())
		http.SetCookie(w, model.GetPlayerCookie(player))

		templ.Handler(
			templates.OpenLobby(
				game.GetPlayers(), game.GetSessionId(), joinUrl)).
			ServeHTTP(w, r)
	}
}

func JoinSession(server *model.ServerState) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var game model.IGameState

		if err := r.ParseForm(); err != nil {
			log.Println("failed to parse session form")
		}

		for _, gameSession := range server.Games {
			if gameSession.GetSessionId() == r.PathValue("session") {
				game = gameSession
				break
			}
		}

		templ.Handler(
			templates.Scaffold(
				templates.JoinLobby(
					game, server.BaseUrl))).
			ServeHTTP(w, r)
	}
}

func JoinedSession(server *model.ServerState) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var game model.IGameState
		player := model.Player()

		if err := r.ParseForm(); err != nil {
			log.Println("failed to parse session form")
			return
		}

		player.SetDisplayName(r.Form.Get("display-name"))

		for _, gameSession := range server.Games {
			if gameSession.GetSessionId() == r.PathValue("session") {
				game = gameSession
				break
			}
		}

		game.Join(player)

		joinUrl := templ.URL(fmt.Sprintf("http://%v/lobby/join/%v", server.BaseUrl, game.GetSessionId()))

		templ.Handler(
			templates.OpenLobby(
				game.GetPlayers(), game.GetSessionId(), joinUrl)).
			ServeHTTP(w, r)

		e := templates.CreatePlayerJoinedEvent(r.Context(), game.GetPlayers(), game.GetSessionId(), joinUrl)
		game.PubEvent(e)
	}
}

func Ready(server *model.ServerState) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		game := GetGame(r, server)
		playerId, err := IsAuthorized(r, game)

		if err != nil {
			log.Printf("failed to get player cookie: %v", err)
			templ.Handler(
				templates.Error("failed to authenticate player."),
			).ServeHTTP(w, r)
			return
		}

		for _, player := range game.GetPlayers() {
			if player.GetId() == *playerId {
				player.SetState(model.Ready)
			}
		}

		//TODO
	}
}
