package grpc

import (
	"context"
	"log"
	"net"

	paymentpb "lnpay/internal/transport/grpc/payment/v1"

	"google.golang.org/grpc"
)

type Server struct {
}

func New() Server {
	return Server{}
}

func (s *Server) Serve(ctx context.Context) error {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	paymentpb.RegisterPaymentServiceServer(grpcServer, &paymentService{})

	startErr := make(chan error)
	go func() {
		startErr <- grpcServer.Serve(lis)
	}()

	select {
	case <-ctx.Done():
		break
	case err := <-startErr:
		if err != nil {
			return err
		}
	}

	log.Println("Shutting down gRPC server...")
	grpcServer.Stop()
	return nil
}
