package store

import (
	"sync"
	"webhook-receiver/internal/model"
)

type MemoryStore struct {
	sync.RWMutex
	events map[string]*model.Event
	subs   []chan *model.Event
}

func (ms *MemoryStore) Subscribe() chan *model.Event {
	ch := make(chan *model.Event, 1)

	ms.Lock()
	ms.subs = append(ms.subs, ch)
	defer ms.Unlock()

	return ch
}

func (ms *MemoryStore) Unsubscribe(ch chan *model.Event) {
	ms.Lock()
	defer ms.Unlock()
	for i, sub := range ms.subs {
		if sub == ch {
			ms.subs = append(ms.subs[:i], ms.subs[i+1:]...)
			close(ch)
			break
		}
	}
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

	for _, ch := range ms.subs {
		select {
		case ch <- event:
		default:
			// 読み手が詰まってたら捨てる（CLIはリアルタイム重視）
		}
	}
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
