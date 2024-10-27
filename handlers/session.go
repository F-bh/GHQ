package handlers

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/f-bh/ghq/model"
	"github.com/f-bh/ghq/templates"
)

func NewSession(server *model.ServerState) func(http.ResponseWriter, *http.Request) {
	player := model.Player()
	game := model.NewGamestate()
	server.Games = append(server.Games, &game)

	game.Players[0] = player

	return func(w http.ResponseWriter, r *http.Request) {
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

		templ.Handler(
			templates.OpenLobby(
				names, game.GetSession())).
			ServeHTTP(w, r)
	}
}
