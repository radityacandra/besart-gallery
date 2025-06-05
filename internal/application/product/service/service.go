package service

import (
	"context"

	"github.com/radityacandra/besart-gallery/api/product"
	"github.com/radityacandra/besart-gallery/internal/application/product/repository"
	"github.com/radityacandra/besart-gallery/internal/application/product/types"
)

type IService interface {
	ProductList(ctx context.Context, input types.ProductListInput) (types.ProductListOutput, error)
	ProductDetail(ctx context.Context, id string) (product.ProductDetailGetResponse, error)
}

type Service struct {
	Repository repository.IRepository
}

func NewService(repository repository.IRepository) *Service {
	return &Service{
		Repository: repository,
	}
}
