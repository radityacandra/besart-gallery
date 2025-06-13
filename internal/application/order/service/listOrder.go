package service

import (
	"context"

	"github.com/radityacandra/besart-gallery/internal/application/order/types"
)

func (s *Service) ListOrder(ctx context.Context, input types.ListOrderInput) (types.ListOrderOutput, error) {
	return s.Repository.GetOrderByUserId(ctx, input)
}
