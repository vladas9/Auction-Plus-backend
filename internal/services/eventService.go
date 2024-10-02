package services

import (
	"sync"

	"github.com/google/uuid"
)

type EventType int

type Event struct {
	Type  EventType `json:"type"`
	Msg   any       `json:"msg"`
	ObjId uuid.UUID
}

const (
	RaisedBidEvent EventType = iota
	AuctionClosedEvent
)

type EventService interface {
	Broadcast(ev *Event)
	Subscribe(ch chan *Event, objId uuid.UUID)
}

func NewEventService() EventService {
	return &eventService{}
}

type eventService struct {
	Chans map[uuid.UUID]([]chan *Event)
	mu    sync.Mutex
}

func (s *eventService) Broadcast(ev *Event) {
	s.mu.Lock()
	chans, ok := s.Chans[ev.ObjId]
	s.mu.Unlock()
	if ok {
		for _, ch := range chans {
			ch <- ev
		}
	}
}

func (s *eventService) Subscribe(ch chan *Event, objId uuid.UUID) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.Chans[objId]; !exists {
		s.Chans[objId] = make([]chan *Event, 0)
	} else {
		s.Chans[objId] = append(s.Chans[objId], ch)
	}
}
