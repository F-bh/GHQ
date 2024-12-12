package templates

import (
	"context"
	"log"
	"strings"

	"github.com/a-h/templ"
	"github.com/f-bh/ghq/model"
)

func CreatePlayerJoinedEvent(ctx context.Context, playerNames []string, sessionId model.GameId, joinUrl templ.SafeURL) model.Event {
	buf := new(strings.Builder)
	err := OpenLobby(playerNames, sessionId, joinUrl).Render(ctx, buf)
	if err != nil {
		log.Printf("failed to create %v event due to:\n %v", model.PlayerJoined, err)
	}

	return model.NewEvent(model.PlayerJoined, buf.String())
}
