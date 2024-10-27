package model

import (
	"math/rand"
	"strconv"
)

type GameId = string

type IGameState interface {
	GetSession() GameId
	GetPlayers() []IPlayer
}

type GameState struct {
	id      GameId
	Tiles   []Tile
	Players [2]IPlayer
}

func (g *GameState) GetSession() GameId {
	return g.id
}

func (g *GameState) GetPlayers() []IPlayer {
	return g.Players[:]
}

func newId() GameId {
	r := rand.Int63()
	return strconv.Itoa(int(r))
}

func NewGamestate() GameState {
	state := GameState{}
	state.id = newId()

	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			tile := DefaultTileData(uint8(x), uint8(y))
			state.Tiles = append(state.Tiles, &tile)
		}
	}

	return state
}
