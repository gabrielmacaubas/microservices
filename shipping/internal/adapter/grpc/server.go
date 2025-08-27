package grpc

import (
	"context"
	"net"

	"github.com/huseyinbabal/microservices/payment/internal/adapters/grpc"
	"github.com/gabrielmacaubas/microservices-proto/golang/shipping"
)

type Server struct {
	proto.UnimplementedShippingServiceServer
	appService application.ShippingService
}

func NewServer(appService application.ShippingService) *Server {
	return &Server{
		appService: appService,
	}
}

func (s *Server) CalculateDeliveryTime(ctx context.Context, req *proto.CalculateDeliveryTimeRequest) (*proto.CalculateDeliveryTimeResponse, error) {
	deliveryTime, err := s.appService.CalculateDeliveryTime(req.Items, req.PurchaseId)
	if err != nil {
		return nil, err
	}
	return &proto.CalculateDeliveryTimeResponse{DeliveryTime: deliveryTime}, nil
}

func (s *Server) Start(address string) error {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	proto.RegisterShippingServiceServer(grpcServer, s)
	return grpcServer.Serve(listener)
}