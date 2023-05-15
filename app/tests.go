package main

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func initTests(app fiber.Router, l *zap.Logger) {
	t := app.Group("/tests")

	t.Get("/debug", func(c *fiber.Ctx) error {
		l.Debug("Debug log")
		return c.SendString("Debug log triggered")
	})
	t.Get("/info", func(c *fiber.Ctx) error {
		l.Info("Info log")
		return c.SendString("Info log triggered")
	})
	t.Get("/warn", func(c *fiber.Ctx) error {
		l.Warn("Warn log")
		return c.SendString("Warn log triggered")
	})
	t.Get("/error", func(c *fiber.Ctx) error {
		l.Error("Error log")
		return c.SendString("Error log triggered")
	})
	t.Get("/various", func(c *fiber.Ctx) error {
		for n := 0; n <= 10; n++ {
			l.Debug("Debug log")
			l.Info("Info log")
			l.Warn("Warn log")
			l.Error("Error log")
		}
		return c.SendString("Various logs triggered")
	})
}
