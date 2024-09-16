package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	m "github.com/vladas9/backend-practice/internal/models"
)

func (c *Controller) BidHandler(w http.ResponseWriter, r *http.Request) error {
	var err error
	bid := &m.BidModel{}
	if err = json.NewDecoder(r.Body).Decode(bid); err != nil {
		return fmt.Errorf("Decoding failed(BidHandler): %s", err)
	}
	if bid.UserId, err = uuid.Parse(r.Header.Get("Authorization")); err != nil {
		return fmt.Errorf("Header parsing failed: %s", err)
	}

}
