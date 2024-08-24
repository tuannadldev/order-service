package grpc

import (
	"context"
	"order-service/internal/order/service"
	"order-service/pkg/logger"
	orderService "order-service/proto/order"
)

type orderGrpcService struct {
	log     logger.Logger
	service *service.OrderService
}

func (o *orderGrpcService) CreateOrder(ctx context.Context, req *orderService.CreateOrderReq) (*orderService.CreateOrderRes, error) {
	//TODO implement me
	panic("implement me")
}

func (o *orderGrpcService) PayOrder(ctx context.Context, req *orderService.PayOrderReq) (*orderService.PayOrderRes, error) {
	//TODO implement me
	panic("implement me")
}

func (o *orderGrpcService) SubmitOrder(ctx context.Context, req *orderService.SubmitOrderReq) (*orderService.SubmitOrderRes, error) {
	//TODO implement me
	panic("implement me")
}

func (o *orderGrpcService) UpdateShoppingCart(ctx context.Context, req *orderService.UpdateShoppingCartReq) (*orderService.UpdateShoppingCartRes, error) {
	//TODO implement me
	panic("implement me")
}

func (o *orderGrpcService) CancelOrder(ctx context.Context, req *orderService.CancelOrderReq) (*orderService.CancelOrderRes, error) {
	//TODO implement me
	panic("implement me")
}

func (o *orderGrpcService) CompleteOrder(ctx context.Context, req *orderService.CompleteOrderReq) (*orderService.CompleteOrderRes, error) {
	//TODO implement me
	panic("implement me")
}

func (o *orderGrpcService) ChangeDeliveryAddress(ctx context.Context, req *orderService.ChangeDeliveryAddressReq) (*orderService.ChangeDeliveryAddressRes, error) {
	//TODO implement me
	panic("implement me")
}

func (o *orderGrpcService) GetOrderByID(ctx context.Context, req *orderService.GetOrderByIDReq) (*orderService.GetOrderByIDRes, error) {
	//TODO implement me
	panic("implement me")
}

func (o *orderGrpcService) Search(ctx context.Context, req *orderService.SearchReq) (*orderService.SearchRes, error) {
	//TODO implement me
	panic("implement me")
}

func InitOrderGrpcService(log logger.Logger, service *service.OrderService) *orderGrpcService {
	return &orderGrpcService{log: log, service: service}
}
