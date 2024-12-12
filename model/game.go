package model

import (
	"math/rand"
	"strconv"
)

type GameId = string

type IGameState interface {
	GetSessionId() GameId
	GetPlayers() []IPlayer
	Join(IPlayer) bool
	PubEvent(Event)
	SubEvents() <-chan Event
}

type GameState struct {
	id       GameId
	Tiles    []Tile
	Players  [2]IPlayer
	eventBus chan Event
}

func (g *GameState) GetSessionId() GameId {
	return g.id
}

func (g *GameState) GetPlayers() []IPlayer {
	return g.Players[:]
}

func (g *GameState) PubEvent(e Event) {
	g.eventBus <- e
}

func (g *GameState) SubEvents() <-chan Event {
	return g.eventBus
}

// returns false if the game is already full
func (g *GameState) Join(p IPlayer) bool {
	if g.Players[1] == nil {
		g.Players[1] = p
		return true
	}

	return false
}

func newId() GameId {
	r := rand.Int63()
	return strconv.Itoa(int(r))
}

func NewGamestate() GameState {
	state := GameState{}
	state.id = newId()
	state.eventBus = make(chan Event, 1)

	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			tile := DefaultTileData(uint8(x), uint8(y))
			state.Tiles = append(state.Tiles, &tile)
		}
	}

	return state
}
