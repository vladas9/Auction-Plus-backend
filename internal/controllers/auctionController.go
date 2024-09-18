package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/shopspring/decimal"
	"github.com/vladas9/backend-practice/internal/dtos"
	s "github.com/vladas9/backend-practice/internal/services"
	"github.com/vladas9/backend-practice/internal/utils"
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

	params := s.AuctionParams{
		Offset:   offset,
		Len:      leangth,
		MaxPrice: maxPrice,
		MinPrice: minPrice,
	}
	problems := params.Validate()
	if problems != nil {
		return &ApiError{Status: 400, ErrorMsg: problems}
	}

	auctionList, itemsList, err := c.service.GetAuctions(params)
	var cards []dtos.AuctionCard
	utils.Logger.Info(auctionList)
	for i, auction := range auctionList {
		cards = append(cards, dtos.MapAuctionCard(i+1, *auction, *itemsList[i]))
	}
	return WriteJSON(w, http.StatusOK, cards)
}
