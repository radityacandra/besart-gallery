package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/radityacandra/besart-gallery/internal/application/order/types"
)

type IRepository interface {
	CreateOrder(ctx context.Context, input types.CreateOrderRepoInput) error
}

type Repository struct {
	Db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Db: db,
	}
}
