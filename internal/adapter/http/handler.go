package handler

import (
	"challenge-goexpert-stress-test/internal/adapter/service"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	tester *service.LoadTester
}

func NewHandler(tester *service.LoadTester) *Handler {
	return &Handler{tester: tester}
}

func (h *Handler) LoadTestHandler(c *fiber.Ctx) error {
	type Request struct {
		URL         string `json:"url"`
		Requests    int    `json:"requests"`
		Concurrency int    `json:"concurrency"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	report := h.tester.RunLoadTest(req.URL, req.Requests, req.Concurrency)

	return c.JSON(report)
}

// Simular um erro 404
func (h *Handler) NotFoundTestHandler(c *fiber.Ctx) error {
	return c.Status(404).JSON(fiber.Map{"error": "Página não encontrada"})
}

// Simular um erro 500
func (h *Handler) ServerErrorTestHandler(c *fiber.Ctx) error {
	return c.Status(500).JSON(fiber.Map{"error": "Erro interno do servidor"})
}
