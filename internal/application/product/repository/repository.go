package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/radityacandra/besart-gallery/internal/application/product/types"
)

type IRepository interface {
	GetProducts(ctx context.Context, input types.GetProductInput) (types.ProductListOutput, error)
	FindProductById(ctx context.Context, id string) (types.FindProductByIdOutput, error)
}

type Repository struct {
	Db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Db: db,
	}
}
