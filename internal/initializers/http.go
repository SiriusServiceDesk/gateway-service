package initializers

import (
	"context"
	"github.com/SiriusServiceDesk/gateway-service/internal/config"
	"github.com/SiriusServiceDesk/gateway-service/internal/middleware"
	"github.com/SiriusServiceDesk/gateway-service/pkg/auth_v1"
	"github.com/SiriusServiceDesk/gateway-service/pkg/logger"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
)

func setupHandlersFromEndpoint(ctx context.Context, mux *runtime.ServeMux, opts []grpc.DialOption, cfg *config.Config) error {
	grpcAddress := cfg.GRPCServer.Address

	if err := auth_v1.RegisterAuthV1HandlerFromEndpoint(ctx, mux, grpcAddress, opts); err != nil {
		return err
	}

	return nil
}

func StartHttpServer(cfg *config.Config, ctx context.Context) error {
	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	if err := setupHandlersFromEndpoint(ctx, mux, opts, cfg); err != nil {
		return err
	}

	logger.Info("http server listening", zap.String("address", cfg.HttpServer.Address))

	return http.ListenAndServe(cfg.HttpServer.Address, middleware.SetupCors(mux))
}
