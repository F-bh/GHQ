package handlers

import (
	"log"
	"net/http"

	"github.com/f-bh/ghq/model"
)

func EventHandler(server *model.ServerState) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		sessionId := r.PathValue("session")

		var game model.IGameState = nil
		for _, g := range server.Games {
			if g.GetSessionId() == sessionId {
				game = g
			}
		}

		if game == nil {
			log.Printf("failed to find sessionId: %v\n", sessionId)
			return
		}

		w.Header().Set("Access-Control-Expose-Headers", "Content-Type")
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

	loop:
		for {
			select {
			case <-r.Context().Done():
				log.Print("SSE socket closed")
				break loop
			case e := <-game.SubEvents():
				err := e.ToSSE(w)
				if err != nil {
					log.Printf("failed to send event: %+v to game: %v\n", e, game.GetSessionId())
					continue
				}
				log.Printf("sent SSE Event: %+v", e)
				w.(http.Flusher).Flush()
			}
		}
	}
}
