package model

import (
	"net/http"
)

type AuthResult = int

const (
	Authorized AuthResult = iota
	MissingCookie
	NoAccess
)

const PlayerCookie = "playerId"

func GetPlayerCookie(p IPlayer) *http.Cookie {
	return &http.Cookie{
		Name:     PlayerCookie,
		Value:    p.GetId(),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	}
}

func IsAuthorized(playerCookie *http.Cookie, g IGameState) (AuthResult, *PlayerId) {
	if playerCookie == nil {
		return MissingCookie, nil
	}

	playerId := playerCookie.Value
	players := g.GetPlayers()
	for _, player := range players {
		if player.GetId() == playerId {
			return Authorized, &playerId
		}
	}

	return NoAccess, nil
}
