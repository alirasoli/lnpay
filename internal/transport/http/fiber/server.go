package fiber

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"lnpay/internal/transport/http"
	"log"
)

type server struct {
}

func New() http.HttpServer {
	return &server{}
}

func (s *server) Serve(ctx context.Context, address string) error {
	app := fiber.New()

	startErr := make(chan error)
	go func() {
		startErr <- app.Listen(address)
	}()

	select {
	case <-ctx.Done():
		break
	case err := <-startErr:
		if err != nil {
			return err
		}
	}

	log.Println("Server shutting down ...")
	return app.Shutdown()
}
