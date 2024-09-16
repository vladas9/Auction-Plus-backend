package controllers

import (
	"net/http"
	"strconv"

	m "github.com/vladas9/backend-practice/internal/models"
	s "github.com/vladas9/backend-practice/internal/services"
)

func (c *Controller) GetAuctions(w http.ResponseWriter, r *http.Request) error {
	if err := r.ParseForm(); err != nil {
		return &ApiError{Status: http.StatusBadRequest, ErrorMsg: "Couldn't parse parameters"}
	}

	offsetStr := r.FormValue("offset")
	leangthStr := r.FormValue("leangth")

	offset, err := atoi(offsetStr)
	if err != nil {
		return &ApiError{Status: http.StatusBadRequest, ErrorMsg: "Invalid offset"}
	}

	leangth, err := atoi(leangthStr)
	if err != nil {
		return &ApiError{Status: http.StatusBadRequest, ErrorMsg: "Invalid leangth"}
	}

	var auctionList []*m.AuctionModel

	auctionList, err = c.service.GetAuctions(s.AuctionParams{
		Offset: offset,
		Len:    leangth,
	})
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, auctionList)
}

func atoi(str string) (int, error) {
	if len(str) == 0 {
		return 0, nil
	}
	return strconv.Atoi(str)
}
