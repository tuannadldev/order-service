package server

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"google.golang.org/grpc/keepalive"
	"order-service/internal/order/service"
	"order-service/pkg/constants"
	orderService "order-service/proto/order"
	"time"

	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"net"
	grpcd "order-service/internal/order/delivery/grpc"
)

const (
	maxConnectionIdle = 10
	gRPCTimeout       = 30
	maxConnectionAge  = 60
	gRPCTime          = 1
)

func (s *server) initGrpcServer() (func() error, *grpc.Server, error) {
	l, err := net.Listen(constants.Tcp, s.cfg.Server.Port)
	if err != nil {
		return nil, nil, errors.Wrap(err, "net.Listen")
	}
	recoveryOpts := []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandler(func(e interface{}) (err error) {
			return nil
		}),
	}

	grpcServer := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: maxConnectionIdle * time.Minute,
			Timeout:           gRPCTimeout * time.Second,
			MaxConnectionAge:  maxConnectionAge * time.Minute,
			Time:              gRPCTime * time.Minute,
		}),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_recovery.UnaryServerInterceptor(recoveryOpts...),
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_recovery.UnaryServerInterceptor(),
		),
		),
	)

	ios := service.InitOrderService(nil, nil)
	orderGrpc := grpcd.InitOrderGrpcService(s.log, ios)
	orderService.RegisterOrderServiceServer(grpcServer, orderGrpc)

	go func() {
		s.log.Infof("GRPC server is listening on port: {%s}", s.cfg.Server.Port)
		s.log.Error(grpcServer.Serve(l))
	}()

	return l.Close, grpcServer, nil
}
