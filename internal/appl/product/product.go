package productappl

import (
	"campusconnect-api/internal/domain/product"
	"campusconnect-api/pkg/utils"
	"context"
)

type productApplicationImpl struct {
	ctx context.Context
}

func (s *productApplicationImpl) validate(e *product.Entity) error {
	if e.Name == "" {
		return product.ErrNameRequired
	}
	if e.Price == 0.0 || e.Price < 0.0 {
		return product.ErrPriceRequired
	}
	return nil
}

func (s *productApplicationImpl) Create(e *product.Entity) (*product.Entity, error) {
	err := s.validate(e)
	if err != nil {
		return nil, err
	}

	ent := &product.Entity{
		ID:    utils.NewIdentity(),
		Name:  e.Name,
		Price: e.Price,
	}

	// productrep.NewRepository()
	return ent, nil
}
