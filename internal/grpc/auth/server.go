package auth

import (
	"context"

	authv1 "github.com/kekaswork/protos/gen/go/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Auth interface {
	Login(
		ctx context.Context,
		email string,
		password string,
		appID int32,
	) (token string, err error)
	RegisterNewUser(
		ctx context.Context,
		email string,
		password string,
	) (userID int64, err error)
	IsAdmin(
		ctx context.Context,
		userID int64,
	) (bool, error)
}

type serverAPI struct {
	authv1.UnimplementedAuthServer
	auth Auth
}

func Register(gRPC *grpc.Server, auth Auth) {
	authv1.RegisterAuthServer(gRPC, &serverAPI{auth: auth})
}

func (s *serverAPI) Login(
	ctx context.Context,
	req *authv1.LoginRequest,
) (*authv1.LoginResponse, error) {
	if err := s.validateLogin(req); err != nil {
		return nil, err
	}

	token, err := s.auth.Login(
		ctx,
		req.GetEmail(),
		req.GetPassword(),
		req.GetAppId(),
	)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &authv1.LoginResponse{
		Token: token,
	}, nil
}

func (s *serverAPI) Register(
	ctx context.Context,
	req *authv1.RegisterRequest,
) (*authv1.RegisterResponse, error) {
	if err := s.validateRegister(req); err != nil {
		return nil, err
	}

	userID, err := s.auth.RegisterNewUser(
		ctx,
		req.GetEmail(),
		req.GetPassword(),
	)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &authv1.RegisterResponse{
		UserId: userID,
	}, nil
}

func (s *serverAPI) IsAdmin(
	ctx context.Context,
	req *authv1.IsAdminRequest,
) (*authv1.IsAdminResponse, error) {
	if err := s.validateIsAdmin(req); err != nil {
		return nil, err
	}

	isAdmin, err := s.auth.IsAdmin(
		ctx,
		req.GetUserId(),
	)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &authv1.IsAdminResponse{
		IsAdmin: isAdmin,
	}, nil
}

func (s *serverAPI) validateLogin(req *authv1.LoginRequest) error {
	if req.GetEmail() == "" {
		return status.Error(codes.InvalidArgument, "email is required")
	}

	if req.GetPassword() == "" {
		return status.Error(codes.InvalidArgument, "password is required")
	}

	if req.GetAppId() == 0 {
		return status.Error(codes.InvalidArgument, "app_id is required")
	}

	return nil
}

func (s *serverAPI) validateRegister(req *authv1.RegisterRequest) error {
	if req.GetEmail() == "" {
		return status.Error(codes.InvalidArgument, "email is required")
	}

	if req.GetPassword() == "" {
		return status.Error(codes.InvalidArgument, "password is required")
	}

	return nil
}

func (s *serverAPI) validateIsAdmin(req *authv1.IsAdminRequest) error {
	if req.GetUserId() == 0 {
		return status.Error(codes.InvalidArgument, "user_id is required")
	}

	return nil
}
