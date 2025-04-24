package responders

import (
	"github.com/edr3x/chi-template/internal/interfaces"
	"go.uber.org/zap"
)

type responders struct {
	logger *zap.Logger
}

func NewResponders() interfaces.Responders {
	return &responders{
		// disablle caller here as we do not need it from global handler
		logger: zap.L().WithOptions(zap.WithCaller(false)),
	}
}
