package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	m "github.com/vladas9/backend-practice/internal/models"
)

func (c *Controller) AddBid(w http.ResponseWriter, r *http.Request) error {
	var err error
	bid := &m.BidModel{}
	if err = json.NewDecoder(r.Body).Decode(bid); err != nil {
		return fmt.Errorf("Decoding failed(BidHandler): %s", err)
	}
	if bid.UserId, err = uuid.Parse(r.Header.Get("Authorization")); err != nil {
		return fmt.Errorf("Header parsing failed: %s", err)
	}

	if err := c.service.NewBid(bid); err != nil {
		return err
	}
	return nil
}

func (c *Controller) BidTable(w http.ResponseWriter, r *http.Request) error {
	var err error
	var limit, offset int
	if limit, err = atoi(r.URL.Query().Get("limit")); err != nil {
		limit = 0
	}
	if offset, err = atoi(r.URL.Query().Get("offset")); err != nil {
		offset = 0
	}
	var userId uuid.UUID
	if userId, err = uuid.Parse(r.Header.Get("Authorization")); err != nil {
		return fmt.Errorf("Header parsing failed: %s", err)
	}

	responce, err := c.service.ShowBidTable(userId, limit, offset)
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, Response{
		"bids_table": responce,
	})
}
