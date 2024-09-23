package models

import "github.com/shopspring/decimal"

// TODO implement switch
type Decimal = decimal.Decimal // usage: decimal.NewFromString()

type Category string

const (
	Electronics Category = "electronics"
	Furniture            = "furniture"
	Cars                 = "cars"
	ArtWork              = "artwork"
	//TODO: To add category elements
)

func IsCategory(str string) bool {
	switch Category(str) {
	case "", Electronics, Furniture, Cars, ArtWork:
		return true
	default:
		return false
	}
}

type Condition string

const (
	New  Condition = "new"
	Used           = "used"
	//TODO: To add condition elements
)

func IsCondition(str string) bool {
	switch Condition(str) {
	case "", New, Used:
		return true
	default:
		return false
	}
}

type Status string

const (
	Shipped    Status = "Shipped"
	OnTheWay          = "on_the_way"
	NotPlanned        = "not_planned"
	//TODO: To add status elements
)

func IsStatus(str string) bool {
	switch Status(str) {
	case Shipped, OnTheWay, NotPlanned:
		return true
	default:
		return false
	}
}
