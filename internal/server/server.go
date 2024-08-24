package server

import (
	"context"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"net/http"
	"order-service/config"
	"order-service/pkg/logger"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type server struct {
	cfg *config.Config
	log logger.Logger
	//im            interceptors.InterceptorManager
	//mw            middlewares.MiddlewareManager
	//os            *service.OrderService
	//v             *validator.Validate
	//mongoClient   *mongo.Client
	//elasticClient *v7.Client
	//echo          *echo.Echo
	//metrics       *metrics.ESMicroserviceMetrics
	http       *http.Server
	doneChanel chan struct{}
}

func InitServer(cfg *config.Config, appLoger logger.Logger) *server {
	return &server{
		cfg:        cfg,
		log:        appLoger,
		http:       nil,
		doneChanel: make(chan struct{}),
	}
}

func (s *server) Run() error {
	s.log.Info(s.cfg)
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()
	// init grpc server
	closeService, grpcServer, err := s.initGrpcServer()

	if err != nil {
		cancel()
		return err
	}
	defer closeService()

	// register healthcheck
	healthpb.RegisterHealthServer(grpcServer, health.NewServer())

	s.healthCheck(ctx)

	<-ctx.Done()
	s.waitDown(time.Duration(s.cfg.Server.WaitShotDownDuration) * time.Second)
	// gracefull shutdown
	grpcServer.GracefulStop()
	<-s.doneChanel
	s.log.Infof("%s server exited properly", s.cfg.Server.ServiceName)

	return nil
}

func (s *server) waitDown(duration time.Duration) {
	go func() {
		time.Sleep(duration)
		s.doneChanel <- struct{}{}
	}()
}

func (s *server) healthCheck(ctx context.Context) {

}
