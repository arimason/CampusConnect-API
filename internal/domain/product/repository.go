package product

type Repository interface {
	Create(e *Entity) error
}
