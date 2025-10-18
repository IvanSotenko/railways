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

func NewStation(name string, ticketPrice decimal.Decimal, seatsNumber int) *Station {
	return &Station{
		Name:        name,
		TicketPrice: ticketPrice,
		SeatsNumber: seatsNumber,
		TicketsSold: 0,
	}
}

func (s *Station) TicketsAvailable() int {
	return s.TicketsSold + s.SeatsNumber
}
