package handlers

import (
	"errors"
	"net/http"

	"github.com/f-bh/ghq/model"
)

func GetGame(r *http.Request, server *model.ServerState) model.IGameState {
	for _, gameSession := range server.Games {
		if gameSession.GetSessionId() == r.PathValue("session") {
			return gameSession
		}
	}
	return nil
}

func IsAuthorized(r *http.Request, game model.IGameState) (*model.PlayerId, error) {
	playerCookie, err := r.Cookie(model.PlayerCookie)
	authorized, playerId := model.IsAuthorized(playerCookie, game)
	if err != nil {
		return nil, err
	}

	switch authorized {
	case model.Authorized:
		return playerId, nil
	case model.MissingCookie:
		return nil, errors.New("failed to authenticate player, missing cookie")
	case model.NoAccess:
		return nil, errors.New("failed to authorize player, no access to lobby")
	default:
		return nil, nil
	}
}
