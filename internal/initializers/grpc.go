package initializers

import (
	"github.com/SiriusServiceDesk/gateway-service/internal/config"
	"github.com/SiriusServiceDesk/gateway-service/internal/handlers"
	"github.com/SiriusServiceDesk/gateway-service/pkg/auth_v1"
	"github.com/SiriusServiceDesk/gateway-service/pkg/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"net"
)

func setupGrpcServers(grpcServer *grpc.Server) {
	auth_v1.RegisterAuthV1Server(grpcServer, handlers.AuthHandler{})
}

func StartGrpcServer(cfg *config.Config) error {
	grpcServer := grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
	)
	reflection.Register(grpcServer)

	setupGrpcServers(grpcServer)

	list, err := net.Listen("tcp", cfg.GRPCServer.Address)
	if err != nil {
		return err
	}

	logger.Info("gRPC server listening", zap.String("address", cfg.GRPCServer.Address))

	return grpcServer.Serve(list)
}
