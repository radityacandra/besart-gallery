package service

import (
	"context"

	"github.com/radityacandra/besart-gallery/api/product"
)

func (s *Service) ProductDetail(ctx context.Context, id string) (product.ProductDetailGetResponse, error) {
	output, err := s.Repository.FindProductById(ctx, id)
	return product.ProductDetailGetResponse(output), err
}
