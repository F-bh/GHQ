package model

import (
	"fmt"
	"net/http"
)

type EventStream = chan Event
type EventName = string

const (
	PlayerJoined EventName = "PlayerJoined"
	Error        EventName = "Error"
)

// Event data is guarenteed to be valid HTMX
type Event struct {
	name EventName
	data string
}

func NewEvent(name EventName, data string) Event {
	return Event{
		name: name,
		data: data,
	}
}

func (e *Event) ToSSE(w http.ResponseWriter) error {
	if _, err := fmt.Fprintf(w, "event: \"%v\"\ndata: \"%v\"\n\n", e.name, e.data); err != nil {
		return err
	}
	return nil
}
