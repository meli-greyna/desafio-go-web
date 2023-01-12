package tickets

import (
	"context"
)

type service struct {
	r Repository
}

type Service interface {
	GetTotalTickets(context.Context, string) (int, error)
	AverageDestination(context.Context, string) (float64, error)
}

func NewService(r Repository) Service {
	return &service{r: r}
}

func (s *service) GetTotalTickets(ctx context.Context, country string) (int, error) {
	tickets, err := s.r.GetTicketByDestination(ctx, country)
	if err != nil {
		return 0, err
	}

	return len(tickets), nil
}

func (s *service) AverageDestination(ctx context.Context, country string) (float64, error) {
	tickets, err := s.r.GetAll(ctx)
	if err != nil {
		return 0.0, err
	}
	byCountry, err := s.r.GetTicketByDestination(ctx, country)
	if err != nil {
		return 0.0, err
	}

	return float64(len(byCountry)) / float64(len(tickets)) * 100, nil
}
