package installment

import (
	"net/http"

	"github.com/fxamauri/api-go-gorm/src/interfaces"
	"github.com/fxamauri/api-go-gorm/src/services"
	"github.com/fxamauri/api-go-gorm/src/structs"
	"github.com/gofiber/fiber/v2"
)

type InstallmentHandler struct {
	router             fiber.Router
	installmentService interfaces.InstallmentService
}

func NewInstallmentHandler(router fiber.Router, serviceContainer services.ServiceContainer) InstallmentHandler {
	return InstallmentHandler{
		router:             router,
		installmentService: serviceContainer.InstallmentService,
	}
}

func (ih InstallmentHandler) SetRoutes() {
	group := ih.router.Group("/installment")

	group.Post("", ih.CreateInstallment)
}

func (ih InstallmentHandler) CreateInstallment(c *fiber.Ctx) error {
	var body structs.Installment

	err := c.BodyParser(&body)
	if err != nil {
		statusCode := http.StatusBadRequest
		return c.Status(statusCode).JSON(structs.Response{
			Data: err.Error(),
			Tag:  http.StatusText(statusCode),
		})
	}

	err = ih.installmentService.Create(body)
	if err != nil {
		statusCode := http.StatusInternalServerError
		return c.Status(statusCode).JSON(structs.Response{
			Data: err.Error(),
			Tag:  http.StatusText(statusCode),
		})
	}

	return c.Status(http.StatusCreated).JSON(structs.Response{
		Data: "installment criado",
	})
}
