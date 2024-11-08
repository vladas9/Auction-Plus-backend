package controllers

import (
	//	"encoding/json"
	//	"fmt"
	//	"io"

	"github.com/google/uuid"
	s "github.com/vladas9/backend-practice/internal/services"
	u "github.com/vladas9/backend-practice/internal/utils"
	"golang.org/x/net/websocket"
)

type EventController interface {
	AuctionEvents(ws *websocket.Conn)
	PrivateAuctEvents(ws *websocket.Conn)
}

type eventController struct {
	EventService s.EventService
}

func (c *eventController) newConn(ws *websocket.Conn) conn {
	return conn{
		ws: ws,
		ch: make(chan *s.Event),
	}
}

func NewEventController(service s.EventService) EventController {
	return &eventController{EventService: service}
}

type conn struct {
	ws *websocket.Conn
	ch chan *s.Event
}

func (c conn) handle() {
	for event := range c.ch {
		websocket.JSON.Send(c.ws, event)
	}
}

func (c *eventController) AuctionEvents(ws *websocket.Conn) {
	auctIdstr := ws.Request().PathValue("id")
	auctId, err := uuid.Parse(auctIdstr)
	if err != nil {
		websocket.JSON.Send(
			ws,
			Response{"error": "auctId not valid: " + err.Error()},
		)
		ws.Close()
		return
	}

	conn := c.newConn(ws)
	c.EventService.Subscribe(conn.ch, auctId)
	u.Logger.Info("ws subscribed to uuid: ", auctId)
	conn.handle()
}

func (*eventController) PrivateAuctEvents(ws *websocket.Conn) {

}
