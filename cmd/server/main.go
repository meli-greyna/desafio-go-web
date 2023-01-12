package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/bootcamp-go/desafio-go-web/docs"
	"github.com/bootcamp-go/desafio-go-web/internal/domain"
	"github.com/bootcamp-go/desafio-go-web/pkg/router"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title MELI Bootcamp API
// @version 1.0
// @description This API handles airline tickets.
// @termsOfService https://fake.url/

// @contact.name API Support
// @contact.url https://real.url

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost/8080
func main() {
	godotenv.Load()
	host := os.Getenv("HOST")

	// Cargo csv.
	list, err := LoadTicketsFromFile("tickets.csv")
	if err != nil {
		panic("Couldn't load tickets")
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })

	docs.SwaggerInfo.Host = host
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router := router.NewRouter(r, list)
	router.MapRoutes()

	if err := r.Run(host); err != nil {
		panic(err)
	}

}

func LoadTicketsFromFile(path string) ([]domain.Ticket, error) {

	var ticketList []domain.Ticket

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}

	csvR := csv.NewReader(file)
	data, err := csvR.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}

	for _, row := range data {
		price, err := strconv.ParseFloat(row[5], 64)
		if err != nil {
			return []domain.Ticket{}, err
		}
		ticketList = append(ticketList, domain.Ticket{
			Id:      row[0],
			Name:    row[1],
			Email:   row[2],
			Country: row[3],
			Time:    row[4],
			Price:   price,
		})
	}

	return ticketList, nil
}
