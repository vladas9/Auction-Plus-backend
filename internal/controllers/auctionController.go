package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/vladas9/backend-practice/internal/dtos"
	"github.com/vladas9/backend-practice/internal/errors"

	s "github.com/vladas9/backend-practice/internal/services"
	u "github.com/vladas9/backend-practice/internal/utils"
)

func (c *Controller) AddAuction(w http.ResponseWriter, r *http.Request) error {
	var err error
	auctionDTO := &dtos.AuctionFull{}

	sellerId, err := u.ExtractUserIDFromToken(r, JwtSecret)
	if err != nil {
		return err
	}

	if err = json.NewDecoder(r.Body).Decode(auctionDTO); err != nil {
		return errors.NotValid(err.Error(), err)
	}

	u.Logger.Info("dto", auctionDTO)
	if auctId, err := c.service.NewAuction(auctionDTO, sellerId); err != nil {
		return errors.Internal(err)
	} else {
		return WriteJSON(w, http.StatusOK, Response{
			"auctionId": auctId})
	}
}

func (c *Controller) GetAuctions(w http.ResponseWriter, r *http.Request) error {
	//if err := r.ParseForm(); err != nil {
	//	return fail(err)
	//}

	offsetStr := r.FormValue("offset")
	leangthStr := r.FormValue("limit")
	minPriceStr := r.FormValue("min_price")
	maxPriceStr := r.FormValue("max_price")
	categoryStr := r.FormValue("category")
	conditionStr := r.FormValue("lotcondition")

	offset, err := atoi(offsetStr)
	if err != nil {
		return errors.NotValid("offset not parsable", err)
	}

	leangth, err := atoi(leangthStr)
	if err != nil {
		return errors.NotValid("limit not parsable", err)
	}

	maxPrice, err := atodec(maxPriceStr)
	if err != nil {
		return errors.NotValid("max_price not parsable", err)
	}
	minPrice, err := atodec(minPriceStr)
	if err != nil {
		return errors.NotValid("min_price not parsable", err)
	}

	params := s.AuctionCardParams{
		Offset:    offset,
		Len:       leangth,
		MaxPrice:  maxPrice,
		MinPrice:  minPrice,
		Category:  categoryStr,
		Condition: conditionStr,
	}

	cards, err := c.service.GetAuctionCards(params)
	if err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, Response{
		"lots": cards,
	})
}

func (c *Controller) GetAuction(w http.ResponseWriter, r *http.Request) error {
	auctId, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		return errors.NotValid("uuid not parsable", err)
	}
	auct, err := c.service.GetFullAuctionById(auctId)
	if err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, Response{
		"auction": auct,
	})
}

func (c *Controller) AuctionTable(w http.ResponseWriter, r *http.Request) error {
	var err error
	params := s.AuctionTableParams{}
	if params.Limit, err = atoi(r.URL.Query().Get("limit")); err != nil {
		return errors.NotValid("limit not parsable", err)
	}
	if params.Offset, err = atoi(r.URL.Query().Get("offset")); err != nil {
		return errors.NotValid("offset not parsable", err)
	}
	params.UserId, err = u.ExtractUserIDFromToken(r, JwtSecret)
	if err != nil {
		return err
	}
	response, err := c.service.GetAuctionTable(params)
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, Response{
		"lots_table": response,
	})

}
