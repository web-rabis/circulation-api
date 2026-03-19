package order

import "errors"

var (
	ErrOrdersTooHigh        = errors.New("orders too high")
	ErrOrdersAlreadyOrdered = errors.New("orders already ordered")
)
