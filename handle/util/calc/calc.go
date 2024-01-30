package calc

import (
	"github.com/shopspring/decimal"
)

// Add returns the sum of two float64 values as a float64 value
func Add(x, y float64) float64 {
	v, _ := decimal.NewFromFloat(x).Add(decimal.NewFromFloat(y)).Float64()
	return v
}

// Sub returns the difference of two float64 values as a float64 value
func Sub(x, y float64) float64 {
	v, _ := decimal.NewFromFloat(x).Sub(decimal.NewFromFloat(y)).Float64()
	return v
}

// Mul returns the product of two float64 values as a float64 value
func Mul(x, y float64) float64 {
	v, _ := decimal.NewFromFloat(x).Mul(decimal.NewFromFloat(y)).Float64()
	return v
}

// Div returns the quotient of two float64 values as a float64 value
func Div(x, y float64) float64 {
	v, _ := decimal.NewFromFloat(x).Div(decimal.NewFromFloat(y)).Float64()
	return v
}

// Num is a wrapper type for big.Float that supports chainable operations
type Num struct {
	value *decimal.Decimal
}

// NewNum returns a new Num with the given float64 value
func NewNum(x float64) *Num {
	v := decimal.NewFromFloat(x)
	return &Num{value: &v}
}

// Add returns a new Num that is the sum of this Num and the given float64 value
func (n *Num) Add(x float64) *Num {
	v := n.value.Add(decimal.NewFromFloat(x))
	return &Num{value: &v}
}

// Sub returns a new Num that is the difference of this Num and the given float64 value
func (n *Num) Sub(x float64) *Num {
	v := n.value.Sub(decimal.NewFromFloat(x))
	return &Num{value: &v}
}

// Mul returns a new Num that is the product of this Num and the given float64 value
func (n *Num) Mul(x float64) *Num {
	v := n.value.Mul(decimal.NewFromFloat(x))
	return &Num{value: &v}
}

// Div returns a new Num that is the quotient of this Num and the given float64 value
func (n *Num) Div(x float64) *Num {
	v := n.value.Div(decimal.NewFromFloat(x))
	return &Num{value: &v}
}

// Round returns a new Num that is the nearest integer to this Num
// If this Num is halfway between two integers, it rounds to the even one
func (n *Num) Round(places int32) *Num {
	v := n.value.Round(places)
	return &Num{value: &v}
}

// Ceil returns a new Num that is the smallest integer greater than or equal to this Num
func (n *Num) Ceil() *Num {
	v := n.value.Ceil()
	return &Num{value: &v}
}

// Floor returns a new Num that is the largest integer less than or equal to this Num
func (n *Num) Floor() *Num {
	v := n.value.Floor()
	return &Num{value: &v}
}

// Float64 returns the float64 value of this Num
func (n *Num) Float64() float64 {
	x, _ := n.value.Float64()
	return x
}

// Int64 returns the int64 representation of this Num
func (n *Num) Int64() int64 {
	return n.value.IntPart()
}

// String returns the string representation of this Num
func (n *Num) String() string {
	return n.value.String()
}
