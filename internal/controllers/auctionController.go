package controllers

import (
	"fmt"
	s "github.com/vladas9/backend-practice/internal/services"
	"net/http"
	"strconv"
)

func (c *Controller) GetAuctions(w http.ResponseWriter, r *http.Request) error {
	fail := func(err error) error {
		return fmt.Errorf("GetAuctions controller: %w", err)
	}
	if err := r.ParseForm(); err != nil {
		return fail(err)
	}

	offsetStr := r.FormValue("offset")
	leangthStr := r.FormValue("limit")

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		return fail(err)
	}

	leangth, err := strconv.Atoi(leangthStr)
	if err != nil {
		return fail(err)
	}

	params := s.AuctionParams{
		Offset: offset,
		Len:    leangth}
	problems := params.Validate()
	if problems != nil {
		return &ApiError{Status: 400, ErrorMsg: problems}
	}

	auctionList, err := c.service.GetAuctions(params)
	return WriteJSON(w, http.StatusOK, auctionList)
}
