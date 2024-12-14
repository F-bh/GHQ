package model

import (
	"math/rand"
	"strconv"
	"sync"
)

type GameId = string

type IGameState interface {
	GetSessionId() GameId
	GetPlayers() []IPlayer
	Join(IPlayer) bool
	PubEvent(Event)
	SubEvents(<-chan bool) <-chan Event
}

type eventBus struct {
	sync.Mutex
	subscribers []Subscription
}

type Subscription struct {
	eventChannel chan<- Event
	done         <-chan bool
}

type GameState struct {
	id      GameId
	Tiles   []Tile
	Players [2]IPlayer
	events  eventBus
}

func (g *GameState) GetSessionId() GameId {
	return g.id
}

func (g *GameState) GetPlayers() []IPlayer {
	return g.Players[:]
}

func (g *GameState) PubEvent(e Event) {
	g.events.Lock()
	defer g.events.Unlock()
	for ix, subscriber := range g.events.subscribers {
		select {
		case <-subscriber.done:
			g.events.subscribers = append(g.events.subscribers[:ix], g.events.subscribers[ix+1:]...)
			close(subscriber.eventChannel)
			continue
		default:
		}

		subscriber.eventChannel <- e
	}
}

func (g *GameState) SubEvents(ch <-chan bool) <-chan Event {
	g.events.Lock()
	defer g.events.Unlock()

	eventChan := make(chan Event, 1)
	g.events.subscribers = append(g.events.subscribers, Subscription{
		eventChannel: eventChan,
		done:         ch,
	})
	return eventChan
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

func NewGamestate() *GameState {
	state := GameState{}
	state.id = newId()
	state.events = eventBus{}

	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			tile := DefaultTileData(uint8(x), uint8(y))
			state.Tiles = append(state.Tiles, &tile)
		}
	}

	return &state
}
