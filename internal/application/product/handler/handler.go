package handler

import (
	"github.com/radityacandra/besart-gallery/internal/application/product/repository"
	"github.com/radityacandra/besart-gallery/internal/application/product/service"
	"github.com/radityacandra/besart-gallery/internal/core"
	"go.uber.org/zap"
)

type Handler struct {
	Service service.IService
	Logger  *zap.Logger
}

func NewHandler(deps *core.Dependency) *Handler {
	repository := repository.NewRepository(deps.DB)
	service := service.NewService(repository)

	return &Handler{
		Service: service,
		Logger:  deps.Logger,
	}
}
