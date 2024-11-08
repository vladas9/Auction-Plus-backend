package utils

import (
	"strconv"

	"github.com/shopspring/decimal"
)

func Atoi(str string) (int, error) {
	if str == "" {
		return 0, nil
	}
	return strconv.Atoi(str)
}

func Atodec(str string) (decimal.Decimal, error) {
	if str == "" {
		return decimal.Zero, nil
	}
	return decimal.NewFromString(str)
}
