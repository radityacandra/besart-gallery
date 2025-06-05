package repository

import (
	"context"

	"github.com/radityacandra/besart-gallery/internal/application/product/types"
)

func (r *Repository) GetProducts(ctx context.Context, input types.GetProductInput) (types.ProductListOutput, error) {
	rows, err := r.Db.QueryxContext(ctx, `
		SELECT id, name, image, original_price, discounted_price, rating FROM public.products
	`)
	if err != nil {
		return types.ProductListOutput{}, err
	}
	defer rows.Close()

	output := types.ProductListOutput{}
	for rows.Next() {
		item := types.ProductListOutputItem{}
		err = rows.StructScan(&item)
		if err != nil {
			return types.ProductListOutput{}, nil
		}

		output = append(output, item)
	}

	return output, nil
}
