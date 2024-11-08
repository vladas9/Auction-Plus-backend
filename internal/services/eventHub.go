package services

import (
	"sync"

	"github.com/google/uuid"
)

type EventType int

type Event struct {
	Type  EventType   `json:"type"`
	Msg   interface{} `json:"msg"`
	ObjId uuid.UUID
}

const (
	RaisedBidEvent EventType = iota
	AuctionClosedEvent
)

type EventHub interface {
	Broadcast(ev *Event)
	Subscribe(ch chan *Event, objId uuid.UUID)
}

func NewEventHub() EventHub {
	return &eventHub{}
}

type eventHub struct {
	Chans map[uuid.UUID]([]chan *Event)
	mu    sync.Mutex
}

func (s *eventHub) Broadcast(ev *Event) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if chans, ok := s.Chans[ev.ObjId]; ok {
		for _, ch := range chans {
			ch <- ev
		}
	}
}

func (s *eventHub) Subscribe(ch chan *Event, objId uuid.UUID) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.Chans[objId]; !exists {
		s.Chans[objId] = make([]chan *Event, 0)
	} else {
		s.Chans[objId] = append(s.Chans[objId], ch)
	}
}
