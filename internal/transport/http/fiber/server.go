package fiber

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"lnpay/internal/service"
	"lnpay/internal/transport/http"
	"log"
)

type server struct {
	app     *fiber.App
	handler *handler
}

type handler struct {
	exchangeService service.Exchange
	paymentService  service.Payment
}

func New(paymentService service.Payment, exchangeService service.Exchange) http.HttpServer {
	return &server{
		app: fiber.New(),
		handler: &handler{
			exchangeService: exchangeService,
			paymentService:  paymentService,
		},
	}
}

func (s *server) Serve(ctx context.Context, address string) error {
	s.setupRouter()

	startErr := make(chan error)
	go func() {
		startErr <- s.app.Listen(address)
	}()

	select {
	case <-ctx.Done():
		break
	case err := <-startErr:
		if err != nil {
			return err
		}
	}

	log.Println("Server is shutting down ...")
	return s.app.Shutdown()
}
