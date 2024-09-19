package controllers

import (
	"net/http"
	"strconv"

	m "github.com/vladas9/backend-practice/internal/models"
	s "github.com/vladas9/backend-practice/internal/services"
	u "github.com/vladas9/backend-practice/internal/utils"
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

func (c *Controller) AuctionTable(w http.ResponseWriter, r *http.Request) error {
	var err error
	var limit, offset int
	if limit, err = atoi(r.URL.Query().Get("limit")); err != nil {
		limit = 0
	}
	if offset, err = atoi(r.URL.Query().Get("offset")); err != nil {
		offset = 0
	}
	userId, err := u.ExtractUserIDFromToken(r, JwtSecret)
	if err != nil {
		return err
	}

	responce, err := c.service.GetAuctionTable(userId, limit, offset)
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, Response{
		"lots_table": responce,
	})

}
