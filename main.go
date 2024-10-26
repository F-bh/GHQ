package main

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/f-bh/ghq/model"
	"github.com/f-bh/ghq/templates"
)

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./static"))

	state := model.GameState{}

	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			tile := model.DefaultTileData(uint8(x), uint8(y))
			state.Tiles = append(state.Tiles, &tile)
		}
	}

	home := templates.Home(state)

	mux.Handle("/", templ.Handler(home))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":3000", mux)
}
