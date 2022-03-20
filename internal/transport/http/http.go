package http

import "context"

type HttpServer interface {
	Serve(ctx context.Context, address string) error
}
