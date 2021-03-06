package models

import "container/list"

type EventType int

const (
	EVENT_JOIN = iota
	EVENT_LEAVE
	EVENT_MESSAGE
)

type Event struct {
	Type EventType
	User string
	Timestamp int
	content string
}

const archiveSize = 20

var archive = list.New()

func NewArchive(event Event)  {
	if archive.Len() > archiveSize {
		archive.Remove(archive.Front())
	}
	archive.PushBack(event)
}

func GetEvents(lastReceived int) []Event {
	events := make([]Event, archive.Len())
	for event := archive.Front(); event != nil; event.Next() {
		e := event.Value.(Event)
		if e.Timestamp > int(lastReceived) {
			events = append(events, e)
		}
	}
	return events
}