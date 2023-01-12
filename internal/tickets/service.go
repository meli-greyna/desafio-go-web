package tickets

import (
	"context"

	"github.com/bootcamp-go/desafio-go-web/internal/domain"
)

type service struct {
	r Repository
}

type Service interface {
	GetTotalTickets(context.Context, string) ([]domain.Ticket, error)
	AverageDestination(context.Context, string) (float64, error)
}

func NewService(r Repository) Service {
	return &service{r: r}
}

func (s *service) GetTotalTickets(ctx context.Context, country string) ([]domain.Ticket, error) {
	tickets, err := s.r.GetTicketByDestination(ctx, country)
	if err != nil {
		return []domain.Ticket{}, err
	}

	return tickets, nil
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
