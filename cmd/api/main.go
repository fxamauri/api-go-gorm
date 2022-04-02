package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fxamauri/api-go-gorm/src/handlers"
	"github.com/fxamauri/api-go-gorm/src/repositories"
	"github.com/fxamauri/api-go-gorm/src/services"
	"github.com/fxamauri/api-go-gorm/src/structs"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	server := fiber.New()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		"localhost",
		"postgres",
		"mysecretpassword",
		"golang-api",
		"5432",
	)
	dialector := postgres.Open(dsn)
	db, err := gorm.Open(dialector)
	if err != nil {
		log.Fatalf("banco nao conectou. Err%v", err.Error())
	}

	db.AutoMigrate(&structs.Installment{})

	containerRepository := repositories.GetRepositories(db)
	servicesContainer := services.GetServices(containerRepository)
	handlers.NewHandlerContainer(server, servicesContainer)

	server.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(structs.Response{
			Data: "pong :)",
		})
	})
	if err := server.Listen(":3001"); err != nil {
		log.Fatalf("aplicacao nao subiu Err=%v", err.Error())
	}
}
