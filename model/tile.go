package model

import "fmt"

type TileState int

const (
	Free TileState = iota
	Occupied
	Moveable
	UnderFire
	WillUnderFire
)

type Tile interface {
	GetColour() *string
	SetColour(c string)
	GetPostion() TwoDPos
}

type TileData struct {
	state  TileState
	colour string
	TwoDPos
}

type TwoDPos struct {
	x,
	y uint8
}

func (pos TwoDPos) ToString() string {
	return fmt.Sprintf("x:%v y:%v", pos.x, pos.y)
}

func (t *TileData) GetColour() *string {
	return &t.colour
}

func (t *TileData) SetColour(c string) {
	t.colour = c
}

func (t *TileData) GetPostion() TwoDPos {
	return t.TwoDPos
}

func DefaultTileData(x, y uint8) TileData {
	return TileData{
		colour: "green",
		state:  Free,
		TwoDPos: TwoDPos{
			x,
			y,
		},
	}
}
