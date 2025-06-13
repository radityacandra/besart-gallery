package service

import (
	"context"

	"github.com/radityacandra/besart-gallery/internal/application/order/types"
)

func (s *Service) DetailOrder(ctx context.Context, userId, orderId string) (types.DetailOrderOutput, error) {
	order, err := s.Repository.FindOrderByIdAndUserId(ctx, orderId, userId)
	if err != nil {
		return types.DetailOrderOutput{}, err
	}

	output := types.DetailOrderOutput{
		Id:        order.Id,
		OrderTime: order.CreatedAt,
		Status:    order.Status,
		ShippingInfo: types.ShippingInfo{
			FullName: order.FullName,
			Address:  order.Address,
			PhoneNo:  order.PhoneNo,
			Notes:    order.Notes,
		},
		OrderItems: []types.OrderItemDetail{},
	}

	for _, orderItem := range order.OrderItems {
		product, _ := s.ProductRepository.FindProductById(ctx, orderItem.ProductId)
		output.OrderItems = append(output.OrderItems, types.OrderItemDetail{
			Id:           orderItem.Id,
			ProductImage: product.Image,
			ProductName:  product.Name,
			Quantity:     orderItem.Quantity,
			ProductPrice: int64(product.DiscountedPrice),
		})
	}

	return output, nil
}
