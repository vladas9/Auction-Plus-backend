package controllers

import (
	"fmt"
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
	leangthStr := r.FormValue("limit")

	println(offsetStr)
	println(leangthStr)

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		return &ApiError{Status: http.StatusBadRequest, ErrorMsg: "Invalid offset"}
	}

	leangth, err := strconv.Atoi(leangthStr)
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
	fmt.Println(auctionList)
	for it := range auctionList {
		fmt.Println(it)
	}

	return WriteJSON(w, http.StatusOK, auctionList)
}
