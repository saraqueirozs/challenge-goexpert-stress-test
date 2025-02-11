package main

import (
	"fmt"
	"os"

	"challenge-goexpert-stress-test/config"
	cliAdapter "challenge-goexpert-stress-test/internal/adapter/cli"
	httpAdapter "challenge-goexpert-stress-test/internal/adapter/http"
	"challenge-goexpert-stress-test/internal/adapter/service"

	"github.com/gofiber/fiber/v2"
)

func main() {
	c := config.LoadConfig()
	tester := service.NewLoadTester()

	if len(os.Args) > 1 {
		cliAdapter.RunCLI(tester)
		return
	}

	app := fiber.New()
	handler := httpAdapter.NewHandler(tester)

	app.Post("/loadtest", handler.LoadTestHandler)
	app.Get("/test-404", handler.NotFoundTestHandler)    // Simular erro 404
	app.Get("/test-500", handler.ServerErrorTestHandler) // Simular erro 500

	fmt.Printf("Servidor rodando na porta %s\n", c.ServerPort)
	app.Listen(":" + c.ServerPort)
}
