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
	if _, err := fmt.Fprintf(w, "event: %s\ndata: %s\n\n", e.name, e.data); err != nil {
		return err
	}
	return nil
}
