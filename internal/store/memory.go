package store

import (
	"sync"
	"webhook-receiver/internal/model"
)

type MemoryStore struct {
	sync.RWMutex
	events map[string]*model.Event
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		events: make(map[string]*model.Event),
	}
}

func (ms *MemoryStore) SaveEvent(event *model.Event) {
	ms.Lock()
	defer ms.Unlock()
	ms.events[event.ID] = event
}

func (ms *MemoryStore) GetAllEvents() []*model.Event {
	ms.RLock()
	defer ms.RUnlock()
	list := make([]*model.Event, 0, len(ms.events))
	for _, event := range ms.events {
		list = append(list, event)
	}
	return list
}

func (ms *MemoryStore) GetEventByID(id string) (*model.Event, bool) {
	ms.RLock()
	defer ms.RUnlock()
	event, exists := ms.events[id]
	return event, exists
}
