package handler

import (
	"github.com/radityacandra/besart-gallery/internal/application/order/repository"
	"github.com/radityacandra/besart-gallery/internal/application/order/service"
	productRepository "github.com/radityacandra/besart-gallery/internal/application/product/repository"
	"github.com/radityacandra/besart-gallery/internal/core"
	"go.uber.org/zap"
)

type Handler struct {
	Service service.IService
	Logger  *zap.Logger
}

func NewHandler(deps *core.Dependency) *Handler {
	productRepo := productRepository.NewRepository(deps.DB)
	orderRepo := repository.NewRepository(deps.DB)
	service := service.NewService(orderRepo, productRepo)

	return &Handler{
		Service: service,
		Logger:  deps.Logger,
	}
}
