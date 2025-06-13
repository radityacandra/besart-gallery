package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/radityacandra/besart-gallery/internal/application/order/types"
)

type IRepository interface {
	CreateOrder(ctx context.Context, input types.CreateOrderRepoInput) error
	FindOrderByIdAndUserId(ctx context.Context, orderId, userId string) (types.CreateOrderRepoInput, error)
	UpdateOrderStatusById(ctx context.Context, input types.UpdateOrderStatusInput) error
	GetOrderByUserId(ctx context.Context, input types.ListOrderInput) (types.ListOrderOutput, error)
}

type Repository struct {
	Db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Db: db,
	}
}
