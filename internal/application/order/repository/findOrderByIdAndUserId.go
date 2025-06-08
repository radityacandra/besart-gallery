package repository

import (
	"context"

	"github.com/radityacandra/besart-gallery/internal/application/order/types"
)

func (r *Repository) FindOrderByIdAndUserId(ctx context.Context, orderId, userId string) (types.CreateOrderRepoInput, error) {
	row := r.Db.QueryRowxContext(ctx, `
		SELECT id, user_id, status, receiver_full_name, receiver_address, receiver_phone_no, shipping_notes, created_at, created_by FROM public.orders WHERE id = $1 AND user_id = $2`, orderId, userId)
	if row.Err() != nil {
		return types.CreateOrderRepoInput{}, row.Err()
	}

	var result types.CreateOrderRepoInput
	if err := row.StructScan(&result); err != nil {
		return types.CreateOrderRepoInput{}, err
	}

	return result, nil
}
