package fiber

import "github.com/gofiber/fiber/v2"

// TODO add validation
func (s *server) setupRouter() {
	s.app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("lnpay api v0.1")
	})

	s.app.Post("/pay", s.handler.pay)

	s.app.Post("/exchange", s.handler.getPriceInSats)
}
