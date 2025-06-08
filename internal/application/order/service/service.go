package service

import (
	"context"

	"github.com/radityacandra/besart-gallery/internal/application/order/repository"
	"github.com/radityacandra/besart-gallery/internal/application/order/types"
	productRepository "github.com/radityacandra/besart-gallery/internal/application/product/repository"
)

type IService interface {
	CreateOrder(ctx context.Context, input types.CreateOrderInput) (string, error)
	UpdateStatus(ctx context.Context, input types.UpdateStatusInput) error
}

type Service struct {
	Repository        repository.IRepository
	ProductRepository productRepository.IRepository
}

func NewService(repo repository.IRepository, productRepo productRepository.IRepository) *Service {
	return &Service{
		Repository:        repo,
		ProductRepository: productRepo,
	}
}
