package handlers

import (
	"net/http"

	"go.uber.org/zap"

	"github.com/edr3x/chi-template/internal/interfaces"
	"github.com/edr3x/chi-template/internal/responders"
	"github.com/edr3x/chi-template/internal/services"
)

type AuthHandler struct {
	res    interfaces.Responders
	logger *zap.Logger
	svc    *services.AuthService
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		svc:    services.NewAuthService(),
		logger: zap.L().Named("AuthHandler"),
		res:    responders.NewResponders(zap.L()),
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	res, err := h.svc.Login(ctx)
	if err != nil {
		h.res.Error(w, r, err)
		return
	}
	h.res.JSON(w, res)
}
