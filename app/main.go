package main

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func main() {
	// Default config that creates a new Fiber instance
	app := fiber.New()

	// Creates a new zap development logging config
	cfg := zap.NewDevelopmentConfig()
	cfg.DisableStacktrace = true
	logger, _ := cfg.Build()
	defer logger.Sync()

	// Log at server start
	logger.Info("The server has started")

	// Creates handlers to add several types of logs
	initTests(app, logger)

	// Serves http requests on 127.0.0.1:8000
	app.Listen(":8000")
}
