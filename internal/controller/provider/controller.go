package provider

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"to-do-planner/internal/config"
	"to-do-planner/internal/service/provider"
)

type ProviderController interface {
	GetProviders(*fiber.Ctx) error
	CreateProvider(*fiber.Ctx) error
}

type providerController struct {
	providerService provider.Service
}

func New(cfg config.Config) ProviderController {
	return &providerController{
		providerService: cfg.Services.Provider,
	}
}

func (ctrl *providerController) CreateProvider(c *fiber.Ctx) error {
	var provider provider.Provider
	if err := c.BodyParser(&provider); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}

	err := ctrl.providerService.Create(c.Context(), provider)
	if err != nil {
		log.Printf("Error creating provider: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error creating provider")
	}
	log.Printf("Provider created: %+v", provider)

	return c.SendStatus(fiber.StatusCreated)
}

func (ctrl *providerController) GetProviders(c *fiber.Ctx) error {
	providers, err := ctrl.providerService.GetProviders(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error fetching providers")
	}

	return c.JSON(providers)
}
