package productrep

import (
	"campusconnect-api/internal/domain/product"
	"context"
	"database/sql"
)

type productRepositoryImpl struct {
	conn *sql.DB
}

func (r *productRepositoryImpl) Create(e *product.Entity) error {
	return nil
}

func NewRepository(ctx context.Context, conn *sql.DB) product.Repository {
	return &productRepositoryImpl{
		conn: conn,
	}
}
