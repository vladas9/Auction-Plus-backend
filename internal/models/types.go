package models

type Dollars struct {
	Exact, Cents uint32
}

type Category string

const (
	Electronics Category = "electronics"
	//TODO: To add category elements
)

type Condition string

const (
	New  Condition = "new"
	Used           = "used"
	//TODO: To add condition elements
)

type Status string

const (
	Shipped Status = "Shipped"
	//TODO: To add status elements
)
