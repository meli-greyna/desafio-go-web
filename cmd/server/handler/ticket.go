package handler

import (
	"net/http"

	"github.com/bootcamp-go/desafio-go-web/internal/tickets"
	"github.com/gin-gonic/gin"
)

type Service struct {
	service tickets.Service
}

func NewService(s tickets.Service) *Service {
	return &Service{
		service: s,
	}
}

// @Summary Get tickets by country
// @Tags Tickets
// @Description Gets the number of tickets bought for a trip to a given country passed by URL parameter
// @Produce json
// @Param token header string true "token"
// @Success 200
// @Router /ticket/getByCountry/:dest [GET]
func (s *Service) GetTicketsByCountry() gin.HandlerFunc {
	return func(c *gin.Context) {

		destination := c.Param("dest")

		tickets, err := s.service.GetTotalTickets(c, destination)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}

		c.JSON(200, tickets)
	}
}

// @Summary Destination percentage
// @Tags Tickets
// @Description Gets the percentage of total tickets bought over the last day with destination of given country
// @Produce json
// @Param token header string true "token"
// @Success 200
// @Router /ticket/getAverage/:dest [GET]
func (s *Service) AverageDestination() gin.HandlerFunc {
	return func(c *gin.Context) {

		destination := c.Param("dest")

		avg, err := s.service.AverageDestination(c, destination)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}

		c.JSON(200, avg)
	}
}
