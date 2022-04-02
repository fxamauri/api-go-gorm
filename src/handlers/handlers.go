package handlers

import (
	"github.com/fxamauri/api-go-gorm/src/handlers/installment"
	"github.com/fxamauri/api-go-gorm/src/services"
	"github.com/gofiber/fiber/v2"
)

func NewHandlerContainer(router fiber.Router, serviceContainer services.ServiceContainer) {
	installment.NewInstallmentHandler(router, serviceContainer).SetRoutes()
}
