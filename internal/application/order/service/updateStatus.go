package service

import (
	"context"
	"time"

	"github.com/radityacandra/besart-gallery/internal/application/order/types"
)

func (s *Service) UpdateStatus(ctx context.Context, input types.UpdateStatusInput) error {
	// find order
	_, err := s.Repository.FindOrderByIdAndUserId(ctx, input.OrderId, input.UserId)
	if err != nil {
		return err
	}

	// update order status
	return s.Repository.UpdateOrderStatusById(ctx, types.UpdateOrderStatusInput{
		OrderId:   input.OrderId,
		Status:    input.Status,
		UpdatedAt: time.Now().UnixMilli(),
		UpdatedBy: input.UserId,
	})
}
