package order

import (
	"context"
	"database/sql"
)

type Repository interface {
	Close()
	PutOrder(ctx context.Context, o Order) error
	GetOrdersForAccount(ctx context.Context, accountID string) ([]Order, error)
}

type postgresRepository struct {
	db *sql.DB
}
