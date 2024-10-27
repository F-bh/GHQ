package model

import (
	"fmt"

	"github.com/google/uuid"
)

type IPlayer interface {
	GetDisplayName() string
	SetDisplayName(string)
	GetId() string
}

type player struct {
	displayName string
	id          string
}

func Player() IPlayer {
	return &player{
		id:          newPlayerId(),
		displayName: "",
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

func (p *player) GetId() string {
	return p.id
}

func newPlayerId() string {
	id, err := uuid.NewRandom()
	if err != nil {
		panic(fmt.Sprintf("failed to generate player uuid!\nerr:%v", err))
	}

	return id.String()
}
