package repository

import (
	"context"

	"github.com/radityacandra/besart-gallery/internal/application/product/types"
)

func (r *Repository) FindProductById(ctx context.Context, id string) (types.FindProductByIdOutput, error) {
	row := r.Db.QueryRowxContext(ctx, `
		SELECT description, dimension, discounted_price, id, image, medium, name, original_price, rating FROM public.products WHERE id = $1
	`, id)
	if row.Err() != nil {
		return types.FindProductByIdOutput{}, row.Err()
	}

	detail := types.FindProductByIdOutput{}
	if err := row.StructScan(&detail); err != nil {
		return types.FindProductByIdOutput{}, err
	}

	return detail, nil
}
