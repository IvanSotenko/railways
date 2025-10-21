package station

import (
	"github.com/shopspring/decimal"
)

type Station struct {
	Name        string
	TicketPrice decimal.Decimal
	SeatsNumber int
	TicketsSold int
}

func NewStation(name string, ticketPrice decimal.Decimal, seatsNumber int) (*Station, error) {
	if ticketPrice.IsNegative() {
		return nil, ErrNegativePrice
	}

	if seatsNumber < 0 {
		return nil, ErrNegativeSeats
	}

	return &Station{
		Name:        name,
		TicketPrice: ticketPrice,
		SeatsNumber: seatsNumber,
		TicketsSold: 0,
	}, nil
}

func (s *Station) TicketsAvailable() int {
	return s.SeatsNumber - s.TicketsSold
}
