package services

import (
	"context"
	"time"

	"go.uber.org/zap"
)

type AuthService struct {
	logger *zap.Logger
}

func NewAuthService() *AuthService {
	return &AuthService{
		logger: zap.L().Named("AuthService"),
	}
}

type LoginResponse struct {
	Token string
}

func (s *AuthService) Login(ctx context.Context) (LoginResponse, error) {
	res := LoginResponse{
		Token: "asdfasdfasdf",
	}

	time.Sleep(2 * time.Second)

	return res, nil
}
