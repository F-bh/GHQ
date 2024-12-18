package model

import (
	"fmt"

	"github.com/google/uuid"
)

type PlayerState = int
type PlayerId = string

const (
	InLobby PlayerState = iota
	Ready
	InGame
	Disconnected
)

type IPlayer interface {
	GetDisplayName() string
	SetDisplayName(string)
	GetId() PlayerId
	GetState() PlayerState
	GetStateString() string
	SetState(PlayerState)
}

type player struct {
	displayName string
	id          PlayerId
	state       PlayerState
}

func Player() IPlayer {
	return &player{
		id:          newPlayerId(),
		displayName: "",
		state:       InLobby,
	}
}

func (p *player) GetDisplayName() string {
	return p.displayName
}

func (p *player) SetDisplayName(in string) {
	p.displayName = in
	if len(in) == 0 {
		p.displayName = "anon"
	}
}

func (p *player) GetId() PlayerId {
	return p.id
}

func (p *player) GetState() PlayerState {
	return p.state
}

func (p *player) GetStateString() string {
	switch p.GetState() {
	case InLobby:
		return "in lobby"
	case Ready:
		return "ready"
	case InGame:
		return "in game"
	case Disconnected:
		return "disconnected"
	default:
		return "unknown"
	}
}

func (p *player) SetState(state PlayerState) {
	p.state = state
}

func newPlayerId() string {
	id, err := uuid.NewRandom()
	if err != nil {
		panic(fmt.Sprintf("failed to generate player uuid!\nerr:%v", err))
	}

	return id.String()
}
