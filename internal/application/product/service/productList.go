package service

import (
	"context"

	"github.com/radityacandra/besart-gallery/internal/application/product/types"
)

func (s *Service) ProductList(ctx context.Context, input types.ProductListInput) (types.ProductListOutput, error) {
	return s.Repository.GetProducts(ctx, types.GetProductInput(input))
}
