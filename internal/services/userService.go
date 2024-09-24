package services

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	m "github.com/vladas9/backend-practice/internal/models"
	r "github.com/vladas9/backend-practice/internal/repository"
	u "github.com/vladas9/backend-practice/internal/utils"
)

func (s *Service) CreateUser(user *m.UserModel) (*m.UserModel, error) {
	var err error

	if user.Password, err = u.HashPassword(user.Password); err != nil {
		return nil, err
	}

	imageUUID := uuid.New().String()

	if err = u.DecodeAndSaveImage(user.Image, ImageDir, imageUUID); err != nil {
		return nil, err
	}

	user.Image = imageUUID

	err = s.store.WithTx(func(stx *r.StoreTx) error {
		user.ID, err = stx.UserRepo().Insert(user)
		return err
	})

	if err != nil {
		return nil, fmt.Errorf("Faled to create user: %v", err.Error())
	}

	return user, nil
}

func (s *Service) CheckUser(user *m.UserModel) (storedUser *m.UserModel, err error) {
	err = s.store.WithTx(func(stx *r.StoreTx) error {
		storedUser, err = stx.UserRepo().GetByProperty("email", user.Email)
		return err
	})

	if err != nil {
		return nil, fmt.Errorf("Failed to find user: %v", err.Error())
	}

	if err = u.CompareHashPassword(user.Password, storedUser.Password); err != nil {
		return nil, err
	}

	return storedUser, nil
}

func (s *Service) GetUserData(id uuid.UUID) (storedUser *m.UserModel, err error) {
	err = s.store.WithTx(func(stx *r.StoreTx) error {
		storedUser, err = stx.UserRepo().GetByProperty("id", id)
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("Failed to get user: %s", err)
	}

	return storedUser, nil
}

func (s *Service) GetUserStats(userID uuid.UUID) (map[string]interface{}, error) {
	boughtCategoryCount := make(map[string]int)
	soldCategoryCount := make(map[string]int)

	boughtPriceRangeCount := make([]int, 5)
	soldPriceRangeCount := make([]int, 5)

	var boughtTransactions []*m.TransactionModel
	var err error
	err = s.store.WithTx(func(stx *r.StoreTx) error {
		boughtTransactions, err = stx.TransactionRepo().GetAll([]r.FilterCondition{
			{Property: "buyer_id", Value: userID},
		})
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get transactions as buyer: %v", err)
	}

	var soldTransactions []*m.TransactionModel
	err = s.store.WithTx(func(stx *r.StoreTx) error {
		soldTransactions, err = stx.TransactionRepo().GetAll([]r.FilterCondition{
			{Property: "seller_id", Value: userID},
		})
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get transactions as seller: %v", err)
	}

	// Define the price ranges
	priceRanges := []struct {
		min float64
		max float64
	}{
		{0, 100},
		{100, 200},
		{200, 300},
		{300, 400},
		{400, 500},
	}

	for _, transaction := range boughtTransactions {
		var auction *m.AuctionModel
		err = s.store.WithTx(func(stx *r.StoreTx) error {
			auction, err = stx.AuctionRepo().GetById(transaction.AuctionId)
			return err
		})
		if err != nil {
			return nil, fmt.Errorf("failed to get auction for transaction %v: %v", transaction.ID, err)
		}

		var item *m.ItemModel
		err = s.store.WithTx(func(stx *r.StoreTx) error {
			item, err = stx.ItemRepo().GetById(auction.ItemId)
			return err
		})
		if err != nil {
			return nil, fmt.Errorf("failed to get item for auction %v: %v", auction.ID, err)
		}

		boughtCategoryCount[string(item.Category)]++

		for i, priceRange := range priceRanges {
			min := decimal.NewFromFloat(priceRange.min)
			max := decimal.NewFromFloat(priceRange.max)

			if transaction.Amount.GreaterThanOrEqual(min) && transaction.Amount.LessThan(max) {
				boughtPriceRangeCount[i]++
				break
			}
		}
	}

	for _, transaction := range soldTransactions {
		var auction *m.AuctionModel
		err = s.store.WithTx(func(stx *r.StoreTx) error {
			auction, err = stx.AuctionRepo().GetById(transaction.AuctionId)
			return err
		})
		if err != nil {
			return nil, fmt.Errorf("failed to get auction for transaction %v: %v", transaction.ID, err)
		}

		var item *m.ItemModel
		err = s.store.WithTx(func(stx *r.StoreTx) error {
			item, err = stx.ItemRepo().GetById(auction.ItemId)
			return err
		})
		if err != nil {
			return nil, fmt.Errorf("failed to get item for auction %v: %v", auction.ID, err)
		}

		soldCategoryCount[string(item.Category)]++

		for i, priceRange := range priceRanges {
			min := decimal.NewFromFloat(priceRange.min)
			max := decimal.NewFromFloat(priceRange.max)

			if transaction.Amount.GreaterThanOrEqual(min) && transaction.Amount.LessThan(max) {
				soldPriceRangeCount[i]++
				break
			}
		}

	}

	result := map[string]interface{}{
		"bought_stats": map[string]interface{}{
			"labels": extractKeys(boughtCategoryCount),
			"data":   extractValues(boughtCategoryCount),
		},
		"sold_stats": map[string]interface{}{
			"labels": extractKeys(soldCategoryCount),
			"data":   extractValues(soldCategoryCount),
		},
		"price_range_stats": map[string]interface{}{
			"labels":      []string{"0-100$", "100-200$", "200-300$", "300-400$", "400-500$"},
			"sold_data":   soldPriceRangeCount,
			"bought_data": boughtPriceRangeCount,
		},
	}

	return result, nil
}

// Helper function to extract keys from a map (categories)
func extractKeys(m map[string]int) []string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

// Helper function to extract values from a map (counts)
func extractValues(m map[string]int) []int {
	values := make([]int, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}
