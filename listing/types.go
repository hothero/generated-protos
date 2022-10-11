package listing

import "github.com/shopspring/decimal"

type Listing struct {
	ID    string
	Title string
	Price decimal.Decimal
}
