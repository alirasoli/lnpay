package fiber

import (
	"github.com/gofiber/fiber/v2"
	"lnpay/internal/entity/constants"
)

type exchangeRequest struct {
	Price float64 `json:"price"`
}

func (h *handler) getPriceInSats(ctx *fiber.Ctx) error {
	var req exchangeRequest
	if err := ctx.BodyParser(&req); err != nil {
		return err
	}
	sats, err := h.exchangeService.GetPriceInSats(ctx.Context(), req.Price, constants.IRT)
	if err != nil {
		return err
	}
	return ctx.JSON(map[string]interface{}{
		"sats": sats,
	})
}
