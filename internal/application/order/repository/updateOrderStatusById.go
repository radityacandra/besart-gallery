package repository

import (
	"context"

	"github.com/radityacandra/besart-gallery/internal/application/order/types"
)

func (r *Repository) UpdateOrderStatusById(ctx context.Context, input types.UpdateOrderStatusInput) error {
	_, err := r.Db.NamedExecContext(ctx, `
		UPDATE public.orders
		SET
			status = :status,
			updated_at = :updated_at,
			updated_by = :updated_by
		WHERE
			id = :id
	`, input)

	return err
}
