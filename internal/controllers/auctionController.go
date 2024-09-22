package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/shopspring/decimal"
	s "github.com/vladas9/backend-practice/internal/services"
	u "github.com/vladas9/backend-practice/internal/utils"
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
	minPriceStr := r.FormValue("min_price")
	maxPriceStr := r.FormValue("max_price")
	categoryStr := r.FormValue("category")
	conditionStr := r.FormValue("lotcondition")

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		return fail(err)
	}

	leangth, err := strconv.Atoi(leangthStr)
	if err != nil {
		return fail(err)
	}

	maxPrice, err := decimal.NewFromString(maxPriceStr)
	if err != nil {
		return fail(err)
	}
	minPrice, err := decimal.NewFromString(minPriceStr)
	if err != nil {
		return fail(err)
	}

	params := s.AuctionCardParams{
		Offset:    offset,
		Len:       leangth,
		MaxPrice:  maxPrice,
		MinPrice:  minPrice,
		Category:  categoryStr,
		Condition: conditionStr,
	}
	problems := params.Validate()
	if problems != nil {
		return &ApiError{Status: 400, ErrorMsg: problems}
	}

	cards, err := c.service.GetAuctionCards(params)
	if err != nil {
		return fmt.Errorf("AuctionController: %w", err)
	}
	return WriteJSON(w, http.StatusOK, Response{
		"lots": cards,
	})
}

func (c *Controller) GetAuction(w http.ResponseWriter, r *Response) error {

	//auctions, err := c.service.GetAuctionById() // use internal/dtos/auctionFull.go and the auction service
	return nil
}

func (c *Controller) AuctionTable(w http.ResponseWriter, r *http.Request) error {
	var err error
	params := s.AuctionTableParams{}
	if params.Limit, err = strconv.Atoi(r.URL.Query().Get("limit")); err != nil {
		return err
	}
	if params.Offset, err = strconv.Atoi(r.URL.Query().Get("offset")); err != nil {
		return err
	}
	params.UserId, err = u.ExtractUserIDFromToken(r, JwtSecret)
	if err != nil {
		return err
	}
	if problems := params.Validate(); problems != nil {
		return &ApiError{Status: 400, ErrorMsg: problems}
	}

	response, err := c.service.GetAuctionTable(params)
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, Response{
		"lots_table": response,
	})

}
