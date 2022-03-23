package fiber

import "github.com/gofiber/fiber/v2"

// TODO add validation for required fields
type PaymentRequest struct {
	Amount      int64  `json:"amount"`
	Currency    string `json:"currency"`
	Description string `json:"description"`
}

type PaymentResponse struct {
	Invoice string `json:"invoice"`
	Hash    string `json:"hash"`
}

func (h *handler) pay(ctx *fiber.Ctx) error {
	var req PaymentRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	hash, invoice, err := h.paymentService.InitPayment(ctx.Context(), req.Amount, "USD", req.Description)
	if err != nil {
		return err
	}
	return ctx.JSON(PaymentResponse{
		Invoice: invoice,
		Hash:    hash,
	})
}
