package handlers

import (
	"context"
	"github.com/SiriusServiceDesk/gateway-service/internal/config"
	"github.com/SiriusServiceDesk/gateway-service/pkg/auth_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AuthHandler struct {
	auth_v1.UnimplementedAuthV1Server
}

func createClientToAuthService(ctx context.Context) (auth_v1.AuthV1Client, error) {
	cfg := config.GetConfig()

	conn, err := grpc.DialContext(ctx, cfg.AuthService.Address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return auth_v1.NewAuthV1Client(conn), nil
}

func (a AuthHandler) Status(ctx context.Context, empty *emptypb.Empty) (*auth_v1.StatusResponse, error) {
	conn, err := createClientToAuthService(ctx)
	if err != nil {
		return nil, err
	}

	response, err := conn.Status(ctx, empty)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (a AuthHandler) Login(ctx context.Context, request *auth_v1.LoginRequest) (*auth_v1.LoginResponse, error) {
	conn, err := createClientToAuthService(ctx)
	if err != nil {
		return nil, err
	}

	response, err := conn.Login(ctx, request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (a AuthHandler) Registration(ctx context.Context, request *auth_v1.RegistrationRequest) (*auth_v1.RegistrationResponse, error) {
	conn, err := createClientToAuthService(ctx)
	if err != nil {
		return nil, err
	}

	response, err := conn.Registration(ctx, request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (a AuthHandler) ConfirmEmail(ctx context.Context, request *auth_v1.ConfirmEmailRequest) (*auth_v1.ConfirmEmailResponse, error) {
	conn, err := createClientToAuthService(ctx)
	if err != nil {
		return nil, err
	}

	response, err := conn.ConfirmEmail(ctx, request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (a AuthHandler) User(ctx context.Context, request *auth_v1.UserRequest) (*auth_v1.UserResponse, error) {
	conn, err := createClientToAuthService(ctx)
	if err != nil {
		return nil, err
	}

	response, err := conn.User(ctx, request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (a AuthHandler) GetUserIdFromToken(ctx context.Context, request *auth_v1.GetUserIdFromTokenRequest) (*auth_v1.GetUserIdFromTokenResponse, error) {
	conn, err := createClientToAuthService(ctx)
	if err != nil {
		return nil, err
	}

	response, err := conn.GetUserIdFromToken(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}
