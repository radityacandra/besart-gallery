package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/radityacandra/besart-gallery/internal/application/order/types"
)

func (s *Service) CreateOrder(ctx context.Context, input types.CreateOrderInput) (string, error) {
	// check if product exist
	for _, item := range input.OrderItems {
		_, err := s.ProductRepository.FindProductById(ctx, item.ProductId)
		if err != nil {
			return "", types.ErrProductNotFound
		}
	}

	orderId := uuid.NewString()
	order := types.CreateOrderRepoInput{
		Id:        orderId,
		UserId:    input.UserId,
		Status:    "open",
		FullName:  input.ShippingInfo.FullName,
		Address:   input.ShippingInfo.Address,
		PhoneNo:   input.ShippingInfo.PhoneNo,
		Notes:     input.ShippingInfo.Notes,
		CreatedAt: time.Now().UnixMilli(),
		CreatedBy: input.UserId,
	}
	for _, item := range input.OrderItems {
		order.OrderItems = append(order.OrderItems, types.OrderItemRepo{
			Id:        uuid.NewString(),
			ProductId: item.ProductId,
			OrderId:   orderId,
			Quantity:  item.Quantity,
			CreatedAt: time.Now().UnixMilli(),
			CreatedBy: input.UserId,
		})
	}

	// save order, order item
	err := s.Repository.CreateOrder(ctx, order)
	if err != nil {
		return "", err
	}

	return orderId, nil
}
