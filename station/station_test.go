package station

import (
	"errors"
	"testing"

	"github.com/shopspring/decimal"
)

func TestNewStation_Success(t *testing.T) {
	price := decimal.NewFromFloat(12.34)
	s, err := NewStation("Central", price, 100)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if s == nil {
		t.Fatalf("expected non-nil Station")
	}
	if s.Name != "Central" {
		t.Errorf("Name = %q; want %q", s.Name, "Central")
	}
	if !s.TicketPrice.Equal(price) {
		t.Errorf("TicketPrice = %s; want %s", s.TicketPrice.String(), price.String())
	}
	if s.SeatsNumber != 100 {
		t.Errorf("SeatsNumber = %d; want %d", s.SeatsNumber, 100)
	}
	if s.TicketsSold != 0 {
		t.Errorf("TicketsSold = %d; want %d", s.TicketsSold, 0)
	}
}

func TestNewStation_NegativePrice(t *testing.T) {
	price := decimal.NewFromFloat(-1)
	_, err := NewStation("BadPrice", price, 10)
	if err == nil {
		t.Fatalf("expected error for negative price, got nil")
	}
	if !errors.Is(err, ErrNegativePrice) {
		t.Fatalf("expected ErrNegativePrice, got %v", err)
	}
}

func TestNewStation_NegativeSeats(t *testing.T) {
	price := decimal.NewFromFloat(1)
	_, err := NewStation("BadSeats", price, -5)
	if err == nil {
		t.Fatalf("expected error for negative seats, got nil")
	}
	if !errors.Is(err, ErrNegativeSeats) {
		t.Fatalf("expected ErrNegativeSeats, got %v", err)
	}
}

func TestTicketsAvailable(t *testing.T) {
	s := &Station{
		Name:        "Small",
		TicketPrice: decimal.NewFromInt(10),
		SeatsNumber: 50,
		TicketsSold: 20,
	}
	want := 30
	if got := s.TicketsAvailable(); got != want {
		t.Fatalf("TicketsAvailable = %d; want %d", got, want)
	}
}
