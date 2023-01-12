package router

import (
	"github.com/bootcamp-go/desafio-go-web/cmd/server/handler"
	"github.com/bootcamp-go/desafio-go-web/internal/domain"
	"github.com/bootcamp-go/desafio-go-web/internal/tickets"
	"github.com/gin-gonic/gin"
)

type router struct {
	e       *gin.Engine
	handler *handler.Service
}

type Router interface {
	MapRoutes()
}

func NewRouter(r *gin.Engine, list []domain.Ticket) Router {
	repo := tickets.NewRepository(list)
	service := tickets.NewService(repo)
	ticketsHandler := handler.NewService(service)

	return &router{
		e:       r,
		handler: ticketsHandler,
	}
}

func (r *router) MapRoutes() {
	r.e.GET("/ticket/getByCountry/:dest", r.handler.GetTicketsByCountry())
	r.e.GET("/ticket/getAverage/:dest", r.handler.AverageDestination())
}
