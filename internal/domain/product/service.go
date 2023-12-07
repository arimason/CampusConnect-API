package product

import "errors"

var (
	ErrIDRequired    = errors.New("id is required")
	ErrNameRequired  = errors.New("name is required")
	ErrPriceRequired = errors.New("price is required")
)

type Service interface {
	Create(e *Entity) error
}
