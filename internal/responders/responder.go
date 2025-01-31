package responders

import (
	"github.com/edr3x/chi-template/internal/interfaces"
	"go.uber.org/zap"
)

type responders struct {
	logger *zap.Logger
}

func NewResponders(logger *zap.Logger) interfaces.Responders {
	return &responders{logger: logger}
}
