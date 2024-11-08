package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/vladas9/backend-practice/internal/errors"
	m "github.com/vladas9/backend-practice/internal/models"
	u "github.com/vladas9/backend-practice/internal/utils"
)

func (c *Controller) AddBid(w http.ResponseWriter, r *http.Request) error {
	var err error
	bid := &m.BidModel{}
	if err = json.NewDecoder(r.Body).Decode(bid); err != nil {
		return errors.NotValid(err.Error(), err)
	}
	bid.UserId, err = u.ExtractUserIDFromToken(r, JwtSecret)
	if err != nil {
		return err
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
		return errors.NotValid("limit not parsable", err)
	}
	if offset, err = strconv.Atoi(r.URL.Query().Get("offset")); err != nil {
		return errors.NotValid("limit not parsable", err)
	}
	userId, err := u.ExtractUserIDFromToken(r, JwtSecret)
	if err != nil {
		return err
	}

	response, err := c.service.GetBidTable(userId, limit, offset)
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, Response{
		"bids_table": response,
	})
}
