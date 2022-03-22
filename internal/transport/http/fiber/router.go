package fiber

import "github.com/gofiber/fiber/v2"

func (s *server) setupRouter() {
	s.app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("lnpay api v0.1")
	})
}
