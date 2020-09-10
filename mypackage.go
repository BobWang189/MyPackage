package mypackage

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func New() {
	fmt.Println("mypackage.New()")
}

func Add(num1, num2 float64) decimal.Decimal {
	d1 := decimal.NewFromFloat(num1).Add(decimal.NewFromFloat(num2))
	return d1
}
