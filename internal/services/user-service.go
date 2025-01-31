package services

import (
	"context"
	"fmt"
	"time"

	"github.com/edr3x/chi-template/internal/entities"
	"github.com/edr3x/chi-template/internal/middlewares"
	"go.uber.org/zap"
)

type UserService struct {
	logger *zap.Logger
}

func NewUserService() *UserService {
	return &UserService{
		logger: zap.L().Named("UserService"),
	}
}

func (s *UserService) GetUser(ctx context.Context, id string) (string, error) {
	time.Sleep(1 * time.Second)

	if id == "69" {
		return "", entities.ErrorBadRequest("NO NO NO!!!")
	}

	mval := middlewares.GetUserData(ctx)

	str := fmt.Sprintf("Got id %v with value from middleware: %v", id, mval)

	return str, nil
}
