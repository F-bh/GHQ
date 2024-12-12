package model

import (
	"fmt"
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

func (e *Event) ToSSE() string {
	return fmt.Sprintf("event: %v\n data: %v\n\n", e.name, e.data)
}
