package grpc

import (
	"context"

	paymentpb "lnpay/internal/transport/grpc/payment/v1"
)

type paymentService struct {
	paymentpb.UnimplementedPaymentServiceServer
}

func (s *paymentService) Pay(ctx context.Context, payment *paymentpb.PaymentRequest) (*paymentpb.PaymentResponse, error) {
	return &paymentpb.PaymentResponse{
		Invoice: "invoice",
		Hash:    "hash",
	}, nil
}
