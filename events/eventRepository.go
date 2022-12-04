package events

import (
	"pravelin/logging"
	"sync"
)

type EventRepository struct {
	sync.Mutex
	storage map[string]*Event
}

func NewEventStore() EventRepository {
	return EventRepository{storage: make(map[string]*Event)}
}

func (eventRepository *EventRepository) GetOrCreateEvent(input EventInput) *Event {
	eventRepository.Lock()
	defer eventRepository.Unlock()

	event, ok := eventRepository.storage[input.SessionId]
	if !ok {
		logging.InfoLog.Printf("Creating new event for session id %s \n", input.SessionId)
		event = NewEvent(input)
		eventRepository.storage[input.SessionId] = event
	}
	return event
}
