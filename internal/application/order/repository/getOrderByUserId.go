package repository

import (
	"context"

	"github.com/radityacandra/besart-gallery/internal/application/order/types"
)

func (r *Repository) GetOrderByUserId(ctx context.Context, input types.ListOrderInput) (types.ListOrderOutput, error) {
	offset := (input.Page - 1) * input.PageSize

	rows, err := r.Db.QueryxContext(ctx, `
		SELECT a.id, a.status, (
			SELECT
				SUM(amount_order_item)
			FROM (
				SELECT
					oi.quantity * COALESCE(p.discounted_price, p.original_price, 0) amount_order_item
				FROM public.order_items oi
					JOIN public.products p
						ON oi.product_id = p.id
				WHERE oi.order_id = a.id
			)
		) total_amount
		FROM public.orders a
		WHERE user_id = $1
		LIMIT $2 OFFSET $3
	`, input.UserId, input.PageSize, offset)
	if err != nil {
		return types.ListOrderOutput{}, err
	}
	defer rows.Close()

	output := []types.OrderOutput{}
	for rows.Next() {
		var order types.OrderOutput
		if err := rows.StructScan(&order); err != nil {
			return types.ListOrderOutput{}, err
		}

		output = append(output, order)
	}

	row := r.Db.QueryRowxContext(ctx, `
		SELECT count(1) total_data FROM public.orders WHERE user_id = $1
	`, input.UserId)
	if row.Err() != nil {
		return types.ListOrderOutput{}, row.Err()
	}
	var totalData int64
	if err := row.Scan(&totalData); err != nil {
		return types.ListOrderOutput{}, err
	}

	return types.ListOrderOutput{
		Data: output,
		Pagination: types.Pagination{
			Page:      input.Page,
			PageSize:  input.PageSize,
			TotalData: totalData,
		},
	}, nil
}
