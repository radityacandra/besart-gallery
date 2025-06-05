package repository

import (
	"context"
	"database/sql"

	"github.com/radityacandra/besart-gallery/internal/application/order/types"
)

func (r *Repository) CreateOrder(ctx context.Context, input types.CreateOrderRepoInput) error {
	tx, err := r.Db.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	_, err = tx.NamedExecContext(ctx, `
		INSERT INTO public.orders(id, user_id, status, receiver_full_name, receiver_address, receiver_phone_no, shipping_notes, created_at, created_by)
		VALUES(:id, :user_id, :status, :receiver_full_name, :receiver_address, :receiver_phone_no, :shipping_notes, :created_at, :created_by)`, &input)
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, item := range input.OrderItems {
		_, err = tx.NamedExecContext(ctx, `
			INSERT INTO public.order_items(id, order_id, product_id, quantity, created_at, created_by)
			VALUES(:id, :order_id, :product_id, :quantity, :created_at, :created_by)`, &item)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}
