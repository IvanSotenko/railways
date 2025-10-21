package station

import "errors"

var ErrNegativePrice = errors.New("ticket price cannot be negative")
var ErrNegativeSeats = errors.New("number of seats cannot be negative")
