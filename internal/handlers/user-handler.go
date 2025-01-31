package handlers

import (
	"net/http"

	"go.uber.org/zap"

	"github.com/edr3x/chi-template/internal/interfaces"
	"github.com/edr3x/chi-template/internal/responders"
	"github.com/edr3x/chi-template/internal/services"
)

type UserHandler struct {
	res    interfaces.Responders
	logger *zap.Logger
	svc    *services.UserService
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		svc:    services.NewUserService(),
		logger: zap.L().Named("UserHandler"),
		res:    responders.NewResponders(zap.L()),
	}
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := r.PathValue("id")
	res, err := h.svc.GetUser(ctx, id)
	if err != nil {
		h.res.Error(w, r, err)
		return
	}

	h.res.JSON(w, res)
}
